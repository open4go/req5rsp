package req

// RefundReq 退款请求
type RefundReq struct {
	// 订单号
	OrderID string `form:"order_id" json:"order_id" xml:"order_id"  binding:"required"`
	// 退款原因一定要写
	Reason string `form:"reason" json:"reason" xml:"reason"  binding:"required"`
}
