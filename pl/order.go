package pl

// payload for order

// OrderPaySuccess 支付成功关键信息
type OrderPaySuccess struct {
	// 渠道
	Channel string `json:"channel"`
	// 三方id
	OpenID string `json:"open_id"`
	// 用户
	UserID string `json:"user_id"`
	// 支付地点
	Location string `json:"location"`
	// 交易id
	TransactionId string `json:"transaction_id"`
	// 商户订单id
	OrderNo string `json:"order_no"`
	// 支付金额字符串
	Amount string `json:"amount"`
	// 支付金额数字
	AmountNum int64 `json:"amount_num"`
	// 时间
	Time int64 `json:"time"`
}
