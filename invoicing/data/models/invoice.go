package models

// Invoice
type Invoice struct {
	ID       string        `json:"id"`
	Customer string        `json:"customer"`
	Items    []InvoiceItem `json:"items"`
}
