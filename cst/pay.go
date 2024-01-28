package cst

// ChannelType 支付渠道类型
type ChannelType int

const (
	// WeChatPay 微信支付
	WeChatPay ChannelType = iota
	// Balance 余额支付
	Balance
	// Kala 卡拉卡支付
	Kala
)
