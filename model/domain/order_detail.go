package domain

type OrderDetail struct {
	OrderId  uint   `gorm:"primaryKey"`
	TicketId uint   `json:"id" form:"id" gorm:"primaryKey"`
	Qty      uint   `json:"qty" form:"qty"`
	Order    Order  `gorm:"foreignKey:OrderId"`
	Ticket   Ticket `gorm:"foreignKey:TicketId"`
}
