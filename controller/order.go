package controller

import (
	"errors"
	"net/http"
	"strconv"
	mid "ticketing/middleware"
	"ticketing/model/domain"
	"ticketing/model/request"
	"ticketing/model/response"
	"ticketing/model/service"
	"time"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	os  service.OrderService
	tc  TicketController
	ods service.OrderDetailsService
}

func NewOrderController(os service.OrderService, ods service.OrderDetailsService, tc TicketController) OrderController {
	return OrderController{
		os:  os,
		tc:  tc,
		ods: ods,
	}
}

func (oc OrderController) Create(c echo.Context) error {
	user_id, _ := mid.ExtractTokenUser(c)

	var order domain.Order
	order.UserId = uint(user_id)
	order.CheckoutDate = time.Now()
	order.Status = "Pending"

	var orderRequest request.NewOrderRequest
	c.Bind(&orderRequest)

	order.TotalPrice = 0
	for _, val := range orderRequest.Tickets {
		ticket, err := oc.tc.ts.Get(val.Id)
		if err != nil {
			return NewErrorResponse(c, http.StatusInternalServerError, err)
		}
		if ticket.ID == 0 {
			return NewErrorResponse(c, http.StatusInternalServerError, errors.New("ticket ID invalid"))
		}

		order.TotalPrice += ticket.Price * val.Qty
		order.OrderDetails = append(order.OrderDetails, domain.OrderDetail{TicketId: ticket.ID, Qty: uint(val.Qty)})
	}

	order, err := oc.os.Save(order)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToOrderResponse(order))
}

func (oc OrderController) GetAll(c echo.Context) error {
	orders, err := oc.os.GetAll()
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for iOrder, order := range orders {
		orderDetails, _ := oc.ods.GetByOrderId(order.ID)
		for _, orderDetail := range orderDetails {
			ticket, _ := oc.tc.ts.Get(orderDetail.TicketId)
			orderDetail.Ticket = ticket
			orders[iOrder].OrderDetails = append(orders[iOrder].OrderDetails, orderDetail)
		}
	}

	return NewSuccessResponse(c, response.ToOrderListResponse(orders))
}

func (oc OrderController) Get(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("order_id"))

	order, err := oc.os.Get(uint(order_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if order.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("Not Found"))
	}

	orderDetails, _ := oc.ods.GetByOrderId(order.ID)
	for _, orderDetail := range orderDetails {
		ticket, _ := oc.tc.ts.Get(orderDetail.TicketId)
		orderDetail.Ticket = ticket
		order.OrderDetails = append(order.OrderDetails, orderDetail)
	}

	return NewSuccessResponse(c, response.ToOrderResponse(order))
}

func (oc OrderController) Delete(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("order_id"))

	order, err := oc.os.Get(uint(order_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if order.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("Not Found"))
	}

	if !oc.ItsMyOrder(c, order) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	order, err = oc.os.Delete(order)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToOrderResponse(order))
}

func (oc OrderController) GetAllByUser(c echo.Context) error {
	tokenUserId, _ := mid.ExtractTokenUser(c)
	user_id, _ := strconv.Atoi(c.Param("user_id"))

	if tokenUserId != uint(user_id) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	orders, err := oc.os.GetAllByUser(uint(user_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return NewSuccessResponse(c, response.ToOrderListResponse(orders))
}

func (oc OrderController) Update(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("order_id"))

	order, err := oc.os.Get(uint(order_id))
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if order.ID == 0 {
		return NewErrorResponse(c, http.StatusNotFound, errors.New("Not Found"))
	}

	if !oc.ItsMyOrder(c, order) {
		return NewErrorResponse(c, http.StatusForbidden, errors.New("Forbidden"))
	}

	var newOrder domain.Order
	c.Bind(&newOrder)

	order.Status = newOrder.Status
	order, err = oc.os.Save(order)
	if err != nil {
		return NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return NewSuccessResponse(c, response.ToOrderResponse(order))
}

func (oc OrderController) ItsMyOrder(c echo.Context, order domain.Order) bool {
	user_id, _ := mid.ExtractTokenUser(c)

	return user_id == order.UserId
}
