package req

import "github.com/open4go/req5rsp/cst"

// PayRequest 支付请求
type PayRequest struct {
	// 状态 （前端根据订单状态选择是否调用支付）
	OrderStatus cst.OrderStatus `form:"status" json:"status" xml:"status"`
	// 订单号
	OrderID string `form:"order_id" json:"order_id" xml:"order_id"  binding:"required"`
	// 订单描述
	Desc string `form:"desc" json:"desc" xml:"desc" binding:"required"`
	// 金额 (单位:分）
	Amount int64 `form:"amount" json:"amount" xml:"amount"  binding:"required"`
	// 订单发生地点 (门店id；渠道：微信；...)
	At string `form:"at" json:"at" xml:"at"`
}

// PayInfo 项目请求体
type PayInfo struct {
	PayMethod  string          `form:"pay_method" json:"pay_method" xml:"pay_method"  binding:"required"`
	PayChannel cst.ChannelType `form:"pay_channel" json:"pay_channel" xml:"pay_channel"`
}

// CloseRequest 关单请求
type CloseRequest struct {
	// 订单号
	OrderID string `form:"order_id" json:"order_id" xml:"order_id"  binding:"required"`
}

// ScanPayRequest 扫码支付
type ScanPayRequest struct {
	// 支付渠道 (由前台根据用户支付偏好，在点餐机上完成勾选）
	Channel cst.ChannelType `form:"channel" json:"channel" xml:"channel" binding:"required"`
	// 订单号 (发起下单后会自动生成订单号并且在支付时自动填入该字段)
	OrderID string `form:"order_id" json:"order_id" xml:"order_id" binding:"required"`
	// 订单描述 （订单自动生成）
	Desc string `form:"desc" json:"desc" xml:"desc" binding:"required"`
	// 金额 (单位:分）（订单自动生成）
	Amount int64 `form:"amount" json:"amount" xml:"amount"  binding:"required"`
	// 订单发生地点 (门店id；渠道：微信；...) （订单自动生成）
	At string `form:"at" json:"at" xml:"at"`
	// 由扫码枪通过扫客户提供的支付二维码生成
	Code string `form:"code" json:"code" xml:"code"`
	// 用户可能会需要积分，因此这里可以是手机号或者账号（一般用户只会记住手机号）
	// 这里不限定死，后台或者增加一个账号类型以便区分
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
}
