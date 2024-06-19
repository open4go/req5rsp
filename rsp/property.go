package rsp

// PropertyPrice 查询规格价格
type PropertyPrice struct {
	// Price 返回总价格
	Price int64 `form:"price" json:"price" xml:"price"`
	// Success 成功查询到的结果个数（例如：下单后，部分属性被下架无法查询到
	Success uint `form:"success" json:"success" xml:"success"`
	// Failed 失败的code结果列表
	Failed []string `form:"failed" json:"failed" xml:"failed"`
}
