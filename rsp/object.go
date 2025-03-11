package rsp

import "github.com/open4go/req5rsp/cst"

// CommonChart 条形图和折线图可用相同的数据结构
type CommonChart struct {
	Type cst.ChartType `json:"type"`
	Head []ChartConf   `json:"head"`
	Data []CommonItem  `json:"data"`
}

// ChartConf
// example: { dataKey: 'a', stroke: '#ff0000', strokeDasharray: '5 5', label: '外卖', name: '外卖' }
type ChartConf struct {
	DataKey         string `json:"key"`
	Stroke          string `json:"stroke"`
	StrokeDasharray string `json:"stroke_dash_array"`
	Label           string `json:"label"`
	Name            string `json:"name"`
}

type CommonItem struct {
	Name string `json:"name"`
	// 以下是支持的数据类型
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
	D int `json:"d"`
	E int `json:"e"`
	F int `json:"f"`
}
