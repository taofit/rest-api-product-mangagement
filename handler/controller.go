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
	conn, err := db.SetDbCon()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	handler := handler{DB: conn}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	{
		routes := r.Group("/product")
		routes.GET("/", handler.GetProducts)
	}

	r.Run()
}
