package req

import "github.com/open4go/xprinter/tp"

type CreateRequest struct {
	PrinterID string         `form:"printer_id" json:"printer_id" xml:"printer_id"  binding:"required"`
	TplID     string         `form:"tpl_id" json:"tpl_id" xml:"tpl_id"`
	Content   tp.ReceiptData `form:"content" json:"content" xml:"content"`
}
