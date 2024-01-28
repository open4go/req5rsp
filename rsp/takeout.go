package rsp

// TakeOutOrder 外卖接口返回
type TakeOutOrder struct {
	Ordercode  string `json:"ordercode"`
	OriginId   string `json:"origin_id"`
	ReturnCode string `json:"return_code"`
	ReturnMsg  string `json:"return_msg"`
}
