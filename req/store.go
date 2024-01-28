package req

// StoreRequest 门店信息
type StoreRequest struct {
	Name      string `form:"name" json:"name" xml:"name"  binding:"required"`
	Mobile    string `form:"mobile" json:"mobile" xml:"mobile"  binding:"required"`
	Longitude string `form:"longitude" json:"longitude" xml:"longitude"  binding:"required"`
	Latitude  string `form:"latitude" json:"latitude" xml:"latitude"  binding:"required"`
	Street    string `form:"street" json:"street" xml:"street"  binding:"required"`
}
