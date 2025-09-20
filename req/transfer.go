package req

// TransferRequest 数据转移请求
type TransferRequest struct {
	// MerchantId 商户号
	MerchantId string `form:"merchantId" json:"merchantId" xml:"merchantId"`
	// Content 商品描述
	TransferDate string `form:"transferDate" json:"transferDate" xml:"transferDate"`
	// Type 类型
	Transferred bool `form:"transferred" json:"transferred" xml:"transferred"`
}
