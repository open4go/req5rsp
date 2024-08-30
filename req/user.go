package req

// UserRequest 用户基本信息请求体
type UserRequest struct {
	Name  string `form:"name" json:"name" xml:"name"  binding:"required"`
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}

// UserPhoneRequest 用户手机号转换
type UserPhoneRequest struct {
	Code string `form:"code" json:"code" xml:"code"  binding:"required"`
}
