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
	// MerchantConf 微信/支付宝/银行商户块。拉卡拉（type=3）允许为空，必填改在 finance-api create。
	MerchantConf *Merchant `form:"merchant_conf" json:"merchant_conf" xml:"merchant_conf"`
	// Lakala 拉卡拉 payjp 密钥。type=3 时由 finance-api 校验 partner_code 与 credential_code。
	Lakala Lakala `form:"lakala" json:"lakala" xml:"lakala"`
	// 支付单位
	Unit cst.PayUnit `form:"unit" json:"unit" xml:"unit"`
	// Enabled 是否启用
	Enabled bool `form:"enabled" json:"enabled" xml:"enabled"`
}

// Lakala 拉卡拉网关密钥。不要复用微信的商户证书或 APIv3。
type Lakala struct {
	PartnerCode    string `form:"partner_code" json:"partner_code" xml:"partner_code"`
	CredentialCode string `form:"credential_code" json:"credential_code" xml:"credential_code"`
	AppID          string `form:"app_id" json:"app_id" xml:"app_id"`
	Env            string `form:"env" json:"env" xml:"env"` // test | prod
	NotifyURL      string `form:"notify_url" json:"notify_url" xml:"notify_url"`
	Redirect       string `form:"redirect" json:"redirect" xml:"redirect"`
	Currency       string `form:"currency" json:"currency" xml:"currency"` // 空则 CNY
	Version        string `form:"version" json:"version" xml:"version"`
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
