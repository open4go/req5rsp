package req

import "github.com/open4go/req5rsp/cst"

// Keys 密钥管理
type Keys struct {
	// Name 密钥名称(例如：微信支付商户)
	Name string `form:"name" json:"name" xml:"name"  binding:"required"`
	// Private 密钥
	Private string `form:"private" json:"private" xml:"private"`
	// Type 密钥类型
	Type string `form:"type" json:"type" xml:"type"  binding:"required"`
	// Type 密钥类型
	MerchantConf Merchant `form:"merchant_conf" json:"merchant_conf" xml:"merchant_conf"  binding:"required"`
	// 支付单位
	Unit cst.PayUnit `form:"unit" json:"unit" xml:"unit"`
	// Enabled 是否启用
	Enabled bool `form:"enabled" json:"enabled" xml:"enabled"`
}

type Merchant struct {
	// 商户id
	MerchantID string `form:"merchant_id" json:"merchant_id" xml:"merchant_id"`
	// 商户证书序列号
	MerchantCertSN string `form:"merchant_cert_sn" json:"merchant_cert_sn" xml:"merchant_cert_sn"`
	// 商户 APIv3 密钥（微信支付），不是小程序 AppSecret
	MerchantAPIKey string `form:"merchant_api_key" json:"merchant_api_key" xml:"merchant_api_key"`
	// 小程序 AppSecret，服务端换 token / 用户身份
	AppSecret string `form:"app_secret" json:"app_secret" xml:"app_secret"`
	// Secret 与 app_secret 同义，兼容只传 secret 的客户端
	Secret string `form:"secret" json:"secret" xml:"secret"`
	// 应用Id
	AppID string `form:"app_id" json:"app_id" xml:"app_id"`
	// 回调地址
	Callback string `form:"callback" json:"callback" xml:"callback"`
}
