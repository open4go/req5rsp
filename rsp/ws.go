package rsp

import "github.com/open4go/req5rsp/cst"

// WSMessage 数据推送服务通用消息体
type WSMessage struct {
	Value     string            `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	Type      cst.WSMessageType `form:"type" json:"type" xml:"type"`
	OrderNews []Order           `form:"order_news,omitempty" json:"order_news,omitempty" xml:"order_news,omitempty"`
	BarLine   BarLineChart      `form:"bar_line,omitempty" json:"bar_line,omitempty" xml:"bar_line,omitempty"`
	Data      []MyData          `form:"data" json:"data" xml:"data"`
}

// BarLineChart 条形图和折线图可用相同的数据结构
type BarLineChart struct {
	Type cst.ChartType `json:"type"`
	Head []LineConf    `json:"head"`
	Data []BarItem     `json:"data"`
}

// LineConf
// example: { dataKey: 'a', stroke: '#ff0000', strokeDasharray: '5 5', label: '外卖', name: '外卖' }
type LineConf struct {
	DataKey         string `json:"key"`
	Stroke          string `json:"stroke"`
	StrokeDasharray string `json:"stroke_dash_array"`
	Label           string `json:"label"`
	Name            string `json:"name"`
}

type BarItem struct {
	Name string `json:"name"`
	// 以下是支持的数据类型
	// 需要与LineConf中的key匹配才会显示
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
	D int `json:"d"`
	E int `json:"e"`
	F int `json:"f"`
}

// Order 通用订单信息
type Order struct {
	// 订单id
	ID string `json:"id"`
	// 时间
	Time int64 `json:"time"`
	// 金额
	Price int `json:"price"`
	// 数量
	Number int `json:"number"`
	// 消费者
	Customer User `json:"customer"`
	// 供应商
	Merchant Merchant `json:"merchant"`
	// 订单状态
	Status int `json:"status"`
}

type User struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	// 头像
	Avatar string `json:"avatar"`
}

type Merchant struct {
	// 名称
	Name string `json:"name"`
	// 商家电话
	Mobile string `json:"mobile"`
	// 商品略图
	Icon string `json:"icon"`
}

// MyData
//
//	 id: any;
//	icon: FC<any>;
//	title?: string;
//	content?: string;
type MyData struct {
	ID      string `json:"id"`
	Icon    string `json:"icon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type WSEvent struct {
	Value any   `json:"value"`
	Type  int   `json:"type"`
	Time  int64 `json:"time"`
}
