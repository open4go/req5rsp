package req

// OrderNotify 订单通知
type OrderNotify struct {
	OrderId string `json:"order_id" binding:"required"` // 订单主键id
	Status  int    `json:"status"`
}
