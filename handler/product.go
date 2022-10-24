package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taofit/rest-api-product-mangagement/entity"
)

func (h handler) GetProducts(c *gin.Context) {
	rows, err := h.DB.Query("select id, name, stock, created_at from product")
	if err != nil {
		response(c, http.StatusNotFound, err.Error())
	}

	defer rows.Close()
	products := []entity.Product{}
	for rows.Next() {
		product := entity.Product{}
		rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Created_at)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func response(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{"message": msg, "statusCode": statusCode})
}
