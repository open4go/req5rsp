package cst

// ChannelType 支付渠道类型
type ChannelType int

const (
	// WeChatPay 微信支付
	WeChatPay ChannelType = iota
	// BalancePay 余额支付
	BalancePay
	// KalaPay 卡拉卡支付
	KalaPay
	// AliPay 支付宝
	AliPay
	// CashPay 现金支付
	CashPay
	// UniPay 银联支付
	UniPay
)
