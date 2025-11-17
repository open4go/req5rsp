package pl

// PayConfig 支付配置信息
// 当支付发起后将填入该消息，后续执行查询任务时不需要再进行查询密钥配置等信息
type PayConfig struct {
	// 应用appid
	AppID string `json:"app_id"`
	// 应用appid
	MchID string `json:"mch_id"`
	// 应用密钥
	AppSecret string `json:"app_secret"`
}

// PayRequestInfo 支付请求
type PayRequestInfo struct {
	// 渠道
	Channel string `json:"channel"`
	// 支付配置信息
	Config PayConfig `json:"config"`
	// 订单创建时间
	CreatedAt int64 `json:"created_at"`
	// 订单id
	OrderNo string `json:"order_no"`
	// 商户id，可以通过该id查询到支付配置信息
	MerchantID string `json:"merchant_id"`
}
