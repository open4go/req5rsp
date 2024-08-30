package rsp

// UserPhoneResponse 用户手机号转换
type UserPhoneResponse struct {
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}
