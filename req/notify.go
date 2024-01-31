package req

import "github.com/open4go/req5rsp/cst"

// OrderNotify 订单通知
type OrderNotify struct {
	OrderId string          `json:"order_id" binding:"required"` // 订单主键id
	Status  cst.OrderStatus `json:"status"`
}
