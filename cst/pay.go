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
	// KalaPay 拉卡拉支付（payjp）。值固定为 2，不要在前面插入新常量。
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
	// LakalaMicropay 拉卡拉被扫（顾客付款码）。
	LakalaMicropay
	// LakalaNativeQR 拉卡拉主扫（屏幕展示聚合码）。
	LakalaNativeQR
	// LakalaMiniPay 拉卡拉小程序。v1 只占位，不调用。
	LakalaMiniPay
)

func (p PayMethod) String() string {
	switch p {
	case WxMiniPay:
		return "微信小程序支付"
	case WxScanQRCode:
		return "微信扫码支付"
	case AliMiniPay:
		return "支付宝小程序支付"
	case AliScanQRCode:
		return "支付宝扫码支付"
	case OfflineCashPay:
		return "线下现金支付"
	case PointsPay:
		return "点券支付（虚拟代币）"
	case BalancePayMethod:
		return "余额支付"
	case LakalaMicropay:
		return "拉卡拉被扫"
	case LakalaNativeQR:
		return "拉卡拉主扫"
	case LakalaMiniPay:
		return "拉卡拉小程序支付"
	default:
		return "其他"
	}
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
