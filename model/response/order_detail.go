package response

import "ticketing/model/domain"

type OrderDetailResponse struct {
	Qty    uint
	Ticket TicketResponse
}

func ToOrderDetailResponse(orderDetail domain.OrderDetail) OrderDetailResponse {
	return OrderDetailResponse{
		Qty:    orderDetail.Qty,
		Ticket: ToTicketResponse(orderDetail.Ticket),
	}
}

func ToOrderDetailListResponse(orderDetails []domain.OrderDetail) []OrderDetailResponse {
	response := []OrderDetailResponse{}
	for _, value := range orderDetails {
		response = append(response, ToOrderDetailResponse(value))
	}
	return response
}
