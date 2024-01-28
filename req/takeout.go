package req

type TakeOut struct {
	PriceToken  string `json:"price_token"`  // 跑腿价格token
	Price       string `json:"price"`        // 外卖跑腿费
	Address     string `json:"address"`      // 地址
	AddressName string `json:"address_name"` // 地址名称
	Lng         string `json:"lng"`          // 经度
	Lat         string `json:"lat"`          // 维度
	Distance    string `json:"distance"`     // 距离
	Name        string `json:"name"`         // 收件人信息
	Phone       string `json:"phone"`        // 收件人手机号
	Remark      string `json:"remark"`       // 备注
}

// TakeOutOrder 下单请求参数
type TakeOutOrder struct {
	OriginId        string `json:"origin_id" binding:"required"`
	PriceToken      string `json:"price_token" binding:"required"`
	OrderPrice      string `json:"total_money"  binding:"required"`
	BalancePaymoney string `json:"need_paymoney"  binding:"required"`
	Receiver        string `json:"receiver"  binding:"required"`
	ReceiverPhone   string `json:"receiver_phone"  binding:"required"`
	Note            string `json:"note"  binding:"required"`
	CallbackUrl     string `json:"callback_url" `
	PushType        string `json:"push_type,default=2"`
	SpecialType     string `json:"special_type"`
	CallmeWithtake  string `json:"callme_withtake"`
}
