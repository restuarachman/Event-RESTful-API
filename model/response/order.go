package response

import (
	"ticketing/model/domain"
	"time"
)

type OrderResponse struct {
	ID           uint
	UserId       uint
	CheckoutDate time.Time
	Status       string
	TotalPrice   int
	OrderDetails []OrderDetailResponse
}

func ToOrderResponse(order domain.Order) OrderResponse {
	return OrderResponse{
		ID:           order.ID,
		UserId:       order.UserId,
		CheckoutDate: order.CheckoutDate,
		Status:       order.Status,
		TotalPrice:   order.TotalPrice,
		OrderDetails: ToOrderDetailListResponse(order.OrderDetails),
	}
}

func ToOrderListResponse(orders []domain.Order) []OrderResponse {
	response := []OrderResponse{}
	for _, value := range orders {
		response = append(response, ToOrderResponse(value))
	}
	return response
}
