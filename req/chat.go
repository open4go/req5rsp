package req

type Chat struct {
	Name       string   `form:"name" json:"name" xml:"name" binding:"required"`
	Title      string   `form:"title" json:"title" xml:"title" binding:"required"`
	Email      string   `json:"email,omitempty" bson:"email,omitempty"`
	Avatar     string   `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Tags       []string `form:"tags" json:"tags" xml:"tags"`
	Gender     string   `form:"gender" json:"gender" xml:"gender" binding:"required"`
	Background string   `form:"background" json:"background" xml:"background" binding:"required"`
	CommandID  string   `form:"command_id" json:"command_id" xml:"command_id" binding:"required"`
}
