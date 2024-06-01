package models

// InvoiceItem
type InvoiceItem struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}
