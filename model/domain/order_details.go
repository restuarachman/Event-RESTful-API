package domain

type OrderDetails struct {
	OrderId  uint `gorm:"primaryKey"`
	TicketId uint `gorm:"primaryKey"`
	Qty      uint
	Order    Order  `gorm:"foreignKey:OrderId"`
	Ticket   Ticket `gorm:"foreignKey:TicketId"`
}
