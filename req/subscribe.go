package req

import (
	"context"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// 先在小程序中完成订阅即可获取模版id
const (
	// PlaceOrderSuccessTp 下单成功通知
	PlaceOrderSuccessTp = "XDH55V-yVYDu4_OhIWvYwXDLX5fvjDBUv2IgY6MDPD4"
	// ReadyToTakeTp 取餐模板
	ReadyToTakeTp = "bI6fgdcs9lK9YtFHz00NVAinh6hX_0bRFaRl2N6Dl1w"
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
	DefaultType MessageType = iota
	// PlaceOrderSuccess 下单成功
	PlaceOrderSuccess
	// PayedDone 支付完成
	PayedDone
	// Produced 制作完成
	Produced
	// ReadyToTake 可以取餐
	ReadyToTake
)

var MessageType2TpID = map[MessageType]string{
	PlaceOrderSuccess: PlaceOrderSuccessTp,
	ReadyToTake:       ReadyToTakeTp,
}

// CreateSubscribeRequest 创建消息订阅
type CreateSubscribeRequest struct {
	// 用户唯一标识
	OpenID string `json:"open_id" bson:"open_id"`
	// 渠道(暂时只支持微信小程序）
	Channel SubscribeChannel `form:"channel" json:"channel" xml:"channel"`
	// Payload 消息主体
	Payload *MessagePayload `form:"payload" json:"payload" xml:"payload"  binding:"required"`
	// 类型
	Type MessageType `form:"type" json:"type" xml:"type"`
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
func MessageRoute(ctx context.Context, req CreateSubscribeRequest) (string, map[string]*subscribe.DataItem) {
	// 获取模板ID
	tpID, ok := MessageType2TpID[req.Type]
	if !ok {
		// 如果找不到匹配的模板ID，可以返回一个默认值或处理错误
		tpID = PlaceOrderSuccessTp
	}

	switch req.Type {
	case PlaceOrderSuccess:
		return tpID, NewMessage(req).PlaceOrderSuccess()
	case ReadyToTake:
		return tpID, NewMessage(req).ReadyToTake()
	default:
		return tpID, NewMessage(req).PlaceOrderSuccess()
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
