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
