package req

// PlaceOrder 处理下单接口
type PlaceOrder struct {
	Remark string  `form:"remark" json:"remark" xml:"remark"`
	Pay    PayInfo `form:"pay" json:"pay" xml:"pay"  binding:"required"`
	// 商品列表
	Items []OrderItem `form:"items" json:"items" xml:"items"  binding:"required"`
	User  UserRequest `form:"user" json:"user" xml:"user"  binding:"required"`
	//Store StoreRequest  `form:"store" json:"store" xml:"store"  binding:"required"`
	StoreID string `form:"storeId" json:"storeId" xml:"storeId"  binding:"required"`
	TableID string `form:"table_id" json:"table_id" xml:"table_id"` // 餐桌id
	// 取餐模式：take-out外卖、in-store堂食、pack：打包自提
	PickupMode string `form:"pickup_mode" json:"pickup_mode" xml:"pickup_mode"`
	// 外卖订单信息
	TakeOut TakeOut `json:"take_out"`
}

// OrderItem 项目请求体
type OrderItem struct {
	ID     string  `form:"id" json:"id" xml:"id"  binding:"required"`
	CateID string  `form:"cate_id" json:"cate_id" xml:"cate_id"  binding:"required"`
	Name   string  `form:"name" json:"name" xml:"name"  binding:"required"`
	Price  float64 `form:"price" json:"price" xml:"price"  binding:"required"`
	Number int     `form:"number" json:"number" xml:"number"  binding:"required"`
	Image  string  `form:"image" json:"image" xml:"image"  binding:"required"`
	// 订单是否有属性
	UseProperty int `form:"use_property" json:"use_property" xml:"use_property"  binding:"required"`
	// 直接显示在订单上的 规格属性
	PropsText string `form:"props_text" json:"props_text" xml:"props_text"  binding:"required"`
	// 规格编号
	Props []int `form:"props" json:"props" xml:"props"  binding:"required"`
}
