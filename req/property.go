package req

// CheckPropertyPrice 查询规格价格
type CheckPropertyPrice struct {
	// PropsSelected 勾选的项
	PropsSelected []CodeWithIDS `form:"props_selected" json:"props_selected" xml:"props_selected"`
}

type CodeWithIDS struct {
	// Code 需要查询的id
	Code string `form:"code" json:"code" xml:"code"  binding:"required"`
	// 选项列表
	IDList []int `form:"id_list" json:"id_list" xml:"id_list"  binding:"required"`
}
