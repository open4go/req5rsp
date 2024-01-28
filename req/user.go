package req

// UserRequest 用户基本信息请求体
type UserRequest struct {
	Name  string `form:"name" json:"name" xml:"name"  binding:"required"`
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}
