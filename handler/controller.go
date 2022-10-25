package handler

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taofit/rest-api-product-mangagement/db"
)

type handler struct {
	DB *sql.DB
}

func RegisterRouter() {
	handler := GetHandler()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	{
		routes := r.Group("/product")
		routes.GET("/", handler.GetProducts)
		routes.GET("/:id", handler.GetProduct)
	}

	r.Run()
}

func GetHandler() handler {
	conn, err := db.SetDbCon()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	
	return handler{DB: conn} 
}
