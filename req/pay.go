package req

import "github.com/open4go/req5rsp/cst"

// PayRequest 支付请求
type PayRequest struct {
	// 订单号
	OrderID string `form:"order_id" json:"order_id" xml:"order_id"  binding:"required"`
	// 订单描述
	Desc string `form:"desc" json:"desc" xml:"desc"`
	// 金额 (单位:分）
	Amount int64 `form:"amount" json:"amount" xml:"amount"  binding:"required"`
}

// PayInfo 项目请求体
type PayInfo struct {
	PayMethod  string          `form:"pay_method" json:"pay_method" xml:"pay_method"  binding:"required"`
	PayChannel cst.ChannelType `form:"pay_channel" json:"pay_channel" xml:"pay_channel"`
}
