package invoice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"irayspace.com/invoicing/data/models"
)

func BindInvoice(g *echo.Group) {
	g.GET("/invoices/:id", getOne)
}

// getOne is get invoice
// @Summary				invoice
// @Description 		get invoice
// @Tags				invoice
// @Accept 				json
// @Produce				json
// @Param 				id path string true "id"
// @Success				200 {object} models.Invoice
// @Router 				/api/v1/invoices/{id} [get]
func getOne(c echo.Context) error {
	invoice := models.Invoice{
		ID:       c.Param("id"),
		Customer: "Иван Иванов Иванович",
		Items: []models.InvoiceItem{
			{Item: "Матрёшка", Quantity: 1},
			{Item: "Борщ", Quantity: 1},
		},
	}
	return c.JSON(http.StatusOK, invoice)
}
