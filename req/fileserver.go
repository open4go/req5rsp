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

// UploadFileInfo 文件rpc 更新绑定请求
type UploadFileInfo struct {
	// 目标id(待更新到目标数据的id）
	// 例如：商品的id
	Id string `form:"id" json:"id" xml:"id"`
	// 更新到该目标数据的具体字段
	// 例如：商品中的image字段
	Field string `form:"field" json:"field" xml:"field"`
	// 图片的具体信息
	// 包含url 等
	Payload Response `form:"payload" json:"payload" xml:"payload"`
}

type Response struct {
	// 文件md5 id
	Id string `form:"id" json:"id" xml:"id"`
	// 文件md5 id
	ContentID string `form:"content_id" json:"content_id" xml:"content_id"`
	// 文件名
	Name string `form:"name" json:"name" xml:"name"`
	// 文件路径
	Format string `form:"format" json:"format" xml:"format"`
	// 文件路径
	Path string `form:"path" json:"path" xml:"path"`
	// 文件大小
	Size uint `form:"size" json:"size" xml:"size"`
	// 文件地址
	Url string `form:"url" json:"url" xml:"url"`
	// 文件大小
	Height int `form:"height" json:"height" xml:"height"`
	// 文件大小
	Width int `form:"width" json:"width" xml:"width"`
}
