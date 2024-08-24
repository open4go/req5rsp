package req

import "github.com/open4go/req5rsp/cst"

// OrderNotify 订单通知
type OrderNotify struct {
	// 三方openID
	OpenId string `json:"open_id" binding:"required"` // 订单主键id
	// 订单id
	OrderId string `json:"order_id"` // 订单主键id
	// 状态
	Status cst.OrderStatus `json:"status"`
}
