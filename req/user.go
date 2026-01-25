package req

import target "github.com/r2day/f9z/member"

// UserRequest 用户基本信息请求体
type UserRequest struct {
	Name  string `form:"name" json:"name" xml:"name"  binding:"required"`
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
}

// UserPhoneRequest 用户手机号转换
type UserPhoneRequest struct {
	Code string `form:"code" json:"code" xml:"code"  binding:"required"`
}

// UserProfileRsp 会员信息
type UserProfileRsp struct {
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID string `json:"id" bson:"_id,omitempty"`
	// 用户来源
	From target.UserFrom `json:"source" bson:"source"`
	// 客户信息
	Identity target.IdentityInfo `json:"identity" bson:"identity"`
	// 登陆审计信息
	Login target.LoginInfo `json:"login" bson:"login"`
	// 个人资产
	Assets target.AssetsInfo `json:"assets" bson:"assets"`
	// 当前邀请人情况
	Invitee target.InviterInfo `json:"invitee" bson:"invitee"`
	// 权限管理
	Permission target.PermissionInfo `json:"permission" bson:"permission"`
	// 收藏门店列表
	FavoriteStore []string `json:"favoriteStore" bson:"favoriteStore"`
}
