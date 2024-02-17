package req

import "github.com/open4go/req5rsp/cst"

/*
	{
	  "command_id": 330,
	  "comment": "Dicta quis aspernatur officiis nam voluptatibus ut est voluptatem quia. Et voluptates illo id et cumque aut tempora ducimus aliquid. Ratione repudiandae itaque modi quidem ad repellendus sapiente laborum est. Natus qui accusamus voluptas dolorem laboriosam illo quidem quia laboriosam.\n \rAut voluptatem dolores provident a provident explicabo. Facilis ex dignissimos quia molestias nulla deserunt est. Eos doloremque id ab animi sequi voluptatem quia quas eos. Suscipit laudantium et modi odio quam molestias minima rem.",
	  "customer_id": 728,
	  "date": "2024-02-15T23:36:06.186Z",
	  "id": 144,
	  "product_id": 27,
	  "rating": 3,
	  "status": "pending"
	}
*/
type Review struct {
	// CommandId 订单id
	CommandId int `json:"command_id"`
	// Comment 评论
	Comment string `json:"comment"`
	// CustomerId 客服id
	CustomerId int `json:"customer_id"`
	// 产品id
	ProductId int `json:"product_id"`
	// 综合评分
	Rating int `json:"rating"`
	// 状态
	Status cst.ReviewStatus `json:"status"`
}
