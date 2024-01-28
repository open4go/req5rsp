package req

import "time"

// PaySUCCESS
// 交易状态，枚举值：
// SUCCESS：支付成功
// REFUND：转入退款
// NOTPAY：未支付
// CLOSED：已关闭
// REVOKED：已撤销（付款码支付）
// USERPAYING：用户支付中（付款码支付）
// PAYERROR：支付失败(其他原因，如银行返回失败)
const (
	PaySUCCESS = "SUCCESS"
)

// WxPayCallback 微信支付回调的加密数据
type WxPayCallback struct {
	TransactionId string `json:"transaction_id"`
	Amount        struct {
		PayerTotal    float64 `json:"payer_total"`
		Total         float64 `json:"total"`
		Currency      string  `json:"currency"`
		PayerCurrency string  `json:"payer_currency"`
	} `json:"amount"`
	Mchid           string `json:"mchid"`
	TradeState      string `json:"trade_state"`
	BankType        string `json:"bank_type"`
	PromotionDetail []struct {
		Amount              int    `json:"amount"`
		WechatpayContribute int    `json:"wechatpay_contribute"`
		CouponId            string `json:"coupon_id"`
		Scope               string `json:"scope"`
		MerchantContribute  int    `json:"merchant_contribute"`
		Name                string `json:"name"`
		OtherContribute     int    `json:"other_contribute"`
		Currency            string `json:"currency"`
		StockId             string `json:"stock_id"`
		GoodsDetail         []struct {
			GoodsRemark    string `json:"goods_remark"`
			Quantity       int    `json:"quantity"`
			DiscountAmount int    `json:"discount_amount"`
			GoodsId        string `json:"goods_id"`
			UnitPrice      int    `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
	SuccessTime time.Time `json:"success_time"`
	Payer       struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	OutTradeNo     string `json:"out_trade_no"`
	AppID          string `json:"AppID"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeType      string `json:"trade_type"`
	Attach         string `json:"attach"`
	SceneInfo      struct {
		DeviceId string `json:"device_id"`
	} `json:"scene_info"`
}

// WxRefundCallback  微信退款回调内容（解密后）
type WxRefundCallback struct {
	Mchid         string `json:"mchid"`
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	RefundId      string `json:"refund_id"`
	OutRefundNo   string `json:"out_refund_no"`

	//RefundStatus 通知的类型：
	//REFUND.SUCCESS：退款成功通知
	//REFUND.ABNORMAL：退款异常通知
	//REFUND.CLOSED：退款关闭通知
	//示例值：REFUND.SUCCESS
	RefundStatus string `json:"refund_status"`

	SuccessTime         time.Time `json:"success_time"`
	UserReceivedAccount string    `json:"user_received_account"`
	Amount              struct {
		Total       int `json:"total"`
		Refund      int `json:"refund"`
		PayerTotal  int `json:"payer_total"`
		PayerRefund int `json:"payer_refund"`
	} `json:"amount"`
}
