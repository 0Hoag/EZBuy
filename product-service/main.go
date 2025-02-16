package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"product-service/connect"
	"product-service/modules/item/transport/service"
)

func main() {
	db, err := connect.InitDatabase()
	if err1 := db.Database("product-service").CreateCollection(context.Background(), "category"); err != nil {
		log.Fatal(err1)
	}
	if err2 := db.Database("product-service").CreateCollection(context.Background(), "product"); err != nil {
		log.Fatal(err2)
	}

	if err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	r := gin.Default()

	scope := r.Group("/v1/api") // scope
	{
		categories := scope.Group("/categories")
		{
			categories.POST("", service.CreateCategories(db))
			categories.GET("/getAll", service.GetAllCategories(db))
			categories.GET("/:id", service.GetCategories(db))
			categories.PATCH("/:id", service.UpdateCategories(db))
			categories.DELETE("/:id", service.DeleteCategories(db))
		}

		products := scope.Group("/products")
		{
			products.POST("", service.CreateProduct(db))
			products.GET("/getAll", service.GetAllProduct(db))
			products.GET("/:id", service.GetProduct(db))
			products.PATCH("/:id", service.UpdateProduct(db))
			products.DELETE("/:id", service.DeleteProduct(db))
		}
	}
	r.Run(":8082")
}
