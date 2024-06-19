package req

// CheckPropertyPrice 查询规格价格
type CheckPropertyPrice struct {
	// Code 需要查询的id
	Code []string `form:"code" json:"code" xml:"code"  binding:"required"`
}
