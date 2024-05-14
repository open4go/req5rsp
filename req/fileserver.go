package req

type BindFile struct {
	// 更新的目标id
	Id string `form:"id" json:"id" xml:"id"`
	// 字段
	Field string `form:"field" json:"field" xml:"field"`
	// 更新的目标url
	Url string `form:"url" json:"url" xml:"url"`
	// 最小尺寸
	MinSize string `form:"min_size" json:"min_size" xml:"min_size"`
	// 最大尺寸
	MaxSize string `form:"max_size" json:"max_size" xml:"max_size"`
}
