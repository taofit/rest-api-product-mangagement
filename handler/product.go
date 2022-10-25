package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taofit/rest-api-product-mangagement/entity"
)

const perPage = 10

func (h handler) GetProducts(c *gin.Context) {
	pageNum, valid := getPagination(c)
	if !valid {
		return
	}
	offSet := (pageNum - 1) * perPage
	rows, err := h.DB.Query(`
		SELECT p.id, p.name, p.stock, p.created_at, string_agg(pp.price || ' ' || pp.currency, ',') price 
		FROM product p
		INNER JOIN product_price pp
		ON p.id = pp.product_id
		GROUP BY p.id
		OFFSET $1 LIMIT $2
	`, offSet, perPage)

	hasErr := hasErr(c, err)
	if hasErr {
		return
	}

	defer rows.Close()
	products := []entity.Product{}
	prices := ""
	for rows.Next() {
		product := entity.Product{}
		rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Created_at, &prices)
		product.Prices = getPriceList(prices)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func (h handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product := entity.Product{}
	row := h.DB.QueryRow(`
		SELECT p.id, p.name, p.stock, p.created_at, string_agg(pp.price || ' ' || pp.currency, ',') price
		FROM product p 
		INNER JOIN product_price pp
		ON p.id = pp.product_id
		GROUP BY p.id
		having p.id = $1`, id)
	prices := ""
	err := row.Scan(&product.Id, &product.Name, &product.Stock, &product.Created_at, &prices)

	hasErr := hasErr(c, err)
	if hasErr {
		return
	}
	product.Prices = getPriceList(prices)
	c.JSON(http.StatusOK, product)
}

func getPriceList(prices string) []entity.Price {
	priceList := strings.Split(prices, ",")
	priceStructList := []entity.Price{}

	for _, price := range priceList {
		priceArr := strings.Split(price, " ")
		if len(priceArr) != 2 {
			continue
		}
		amount := priceArr[0]
		currency := priceArr[1]
		priceStruct := entity.Price{Amount: amount, Currency: currency}
		priceStructList = append(priceStructList, priceStruct)
	}
	return priceStructList
}

func getPagination(c *gin.Context) (pageNum int, valid bool) {
	page := c.Query("page")
	var err error
	if page != "" {
		pageNum, err = strconv.Atoi(page)
		if err != nil {
			log.Println("Missing page id")
			errorResponse(c, http.StatusNotFound, err.Error())
			return pageNum, false
		}
	}
	if pageNum < 1 {
		eMsg := "Page number cannot be less than 1"
		log.Println(eMsg)
		errorResponse(c, http.StatusNotFound, eMsg)
		return pageNum, false
	}
	return pageNum, true
}

func hasErr(c *gin.Context, err error) bool {
	if err != nil {
		var eMsg string
		if err == sql.ErrNoRows {
			eMsg = "No row record found in the product table"
			log.Println(eMsg)
		}
		if eMsg == "" {
			eMsg = err.Error()
		}
		errorResponse(c, http.StatusNotFound, eMsg)
		return true
	}

	return false
}

func errorResponse(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{"message": msg, "statusCode": statusCode})
}
