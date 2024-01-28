package rsp

// WxPayPrepare 微信预支付响应
type WxPayPrepare struct {
	Appid     string `json:"appid"`
	NonceStr  string `json:"noncestr"`
	Package   string `json:"package"`
	PartnerID string `json:"partnerid"`
	PrepayID  string `json:"prepayid"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	// 订单id
	OrderID string `json:"order_id"`
}
