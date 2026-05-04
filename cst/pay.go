package cst

// ChannelType 支付渠道
// 支付渠道	实际完成资金划转的通道或服务商，
// 负责对接银行、清算机构、第三方支付平台	微信支付、支付宝、银联云闪付、PayPal、Stripe、银行直连
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

// PayMethod 支付方式 用户看到的支付工具或支付形式，是用户在结账时选择的交互选项
// 余额支付、扫码支付、现金支付、银行卡支付、信用卡、分期付款
type PayMethod int

const (
	// WxMiniPay 微信小程序支付
	WxMiniPay PayMethod = iota
	// WxScanQRCode 微信扫码支付
	WxScanQRCode
	// AliMiniPay 支付宝小程序支付
	AliMiniPay
	// AliScanQRCode 支付宝扫码支付
	AliScanQRCode
	OfflineCashPay // 线下现金支付
	PointsPay      // 点券支付（虚拟代币）
	BalancePayMethod
)

func (p PayMethod) String() string {
	return [...]string{"微信小程序支付", "微信扫码支付", "支付宝小程序支付", "支付宝扫码支付", "线下现金支付", "点券支付（虚拟代币", "余额支付"}[p]
}

type PayStatus int

const (
	// UnPaid 未支付
	UnPaid PayStatus = iota
	// PaidDone 已支付
	PaidDone
	// Refunding 退款中
	Refunding
	// Refunded 已退款
	Refunded
)

func (p PayStatus) String() string {
	return [...]string{"未支付", "已支付", "退款中", "已退款"}[p]
}

type PayUnit int

const (
	// PayByFen 以分支付
	PayByFen PayUnit = 1
	// PayByYuan 以元支付
	PayByYuan PayUnit = 100
)
