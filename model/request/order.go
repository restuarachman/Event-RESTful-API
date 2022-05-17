package request

type OrderRequest struct {
	Id  uint
	Qty int
}

type NewOrderRequest struct {
	Tickets []OrderRequest
}
