package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	Id uint
	InvoiceHeaderId uint
	ProductId uint
	CreateAt time.Time
	UpdateAt time.Time
}
