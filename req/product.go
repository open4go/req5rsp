package req

// SalesCount 销售数量统计更新
type SalesCount struct {
	// SalesInfo 销售列表
	SalesInfo []SalesInfo `form:"sales_info" json:"sales_info" xml:"sales_info"`
}

type SalesInfo struct {
	// ID 需要查询的id
	ID string `form:"id" json:"id" xml:"id"  binding:"required"`
	// 销售数量
	Sales int `form:"sales" json:"sales" xml:"sales"  binding:"required"`
}
