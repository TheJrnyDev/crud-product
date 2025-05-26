package routes

import (
	"crud-product-bck/handlers"
	"crud-product-bck/services"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, client *mongo.Client) {
	// Define API version
	v1 := e.Group("/api/v1")

	// Products routes
	initProductRoutes(v1, client)

}

func initProductRoutes(v1 *echo.Group, client *mongo.Client) {
	// Initialize services with client and database name
	productService := services.NewProductService(client, "crud-product")

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)

	// Define product routes
	v1.GET("/products", productHandler.GetProducts)
	v1.POST("/product", productHandler.CreateProduct)
	v1.DELETE("/product", productHandler.DeleteProduct)
	// prodGroup.PUT("/:id", productHandler.UpdateProduct)
}
