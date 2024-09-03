package rsp

// UserPhoneResponse 用户手机号转换
type UserPhoneResponse struct {
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}

// InviteResponse 用户手机号转换
type InviteResponse struct {
	Code string `form:"code" json:"code" xml:"code"  binding:"required"`
}
