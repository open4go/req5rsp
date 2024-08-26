package req

import (
	"context"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

type SubscribeChannel uint

const (
	WxMini SubscribeChannel = iota
	AlipayMini
	DouYin
)

// MessageType 与模版一一对应
type MessageType uint

const (
	// PlaceOrderSuccess 下单成功
	PlaceOrderSuccess MessageType = iota
	// PayedDone 支付完成
	PayedDone
	// Produced 制作完成
	Produced
	// ReadyToTake 可以取餐
	ReadyToTake
)

type CreateSubscribeRequest struct {
	// 渠道
	Channel SubscribeChannel `form:"channel" json:"channel" xml:"channel"  binding:"required"`
	// Payload 消息主体
	Payload *MessagePayload `form:"mobile" json:"mobile" xml:"mobile"  binding:"required"`
	// 类型
	Type MessageType `form:"type" json:"type" xml:"type"  binding:"required"`
}

// MessagePayload https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-message-management/subscribe-message/sendMessage.html
// 标准信息模版
type MessagePayload struct {
	// Name 名称
	Name string `json:"name" bson:"name"`
	// Address 地址
	Address string `json:"address" bson:"address"`
	// Thing 事件:
	Thing string `json:"thing" bson:"thing"`
	// 金额
	Amount string `json:"amount" bson:"amount"`
	// CharacterString 32位以内数字、字母或符号 例如：订单号，排队号等
	CharacterString string `json:"character_string" bson:"character_string"`
	// 时间日期：例如： 2024-08-24 23:23:23 （可能会有其他的时间）
	Date string `json:"date" bson:"date"`
}

func NewMessage(req CreateSubscribeRequest) *MessagePayload {
	return req.Payload
}

// MessageRoute 根据消息类型不同，路由到不同的消息转换
func MessageRoute(ctx context.Context, req CreateSubscribeRequest) map[string]*subscribe.DataItem {

	switch req.Type {
	case PlaceOrderSuccess:
		return NewMessage(req).PlaceOrderSuccess()
	case ReadyToTake:
		return NewMessage(req).ReadyToTake()
	default:
		return NewMessage(req).PlaceOrderSuccess()
	}
}

// PlaceOrderSuccess  构建不同的数据场景
func (payload *MessagePayload) PlaceOrderSuccess() map[string]*subscribe.DataItem {

	return map[string]*subscribe.DataItem{
		"thing5":   {Value: payload.Thing},
		"amount12": {Value: payload.Amount},
		"thing7":   {Value: payload.Thing},
		"date4":    {Value: payload.Date},
	}
}

// ReadyToTake  构建不同的数据场景
// 订单状态:{{phrase16.DATA}}
// 商品名:{{thing6.DATA}}
// 订单金额:{{amount13.DATA}}
// 餐厅名称:{{thing17.DATA}}
// 取餐时间:{{time21.DATA}}
func (payload *MessagePayload) ReadyToTake() map[string]*subscribe.DataItem {
	return map[string]*subscribe.DataItem{
		"phrase16": {Value: payload.Thing},
		"thing6":   {Value: payload.Name},
		"amount12": {Value: payload.Amount},
		"thing7":   {Value: payload.Address},
		"date4":    {Value: payload.Date},
	}
}
