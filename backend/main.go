package main

import (
	config "crud-product-bck/config"
	route "crud-product-bck/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database connection
	client := config.InitDatabase()

	// Initialize Echo framework
	e := echo.New()

	// Allow CORS for all origins
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if c.Request().Method == echo.OPTIONS {
				return c.NoContent(204) // No Content for preflight requests
			}
			return next(c)
		}
	}))

	// Initialize routes
	route.InitRoutes(e, client)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))

	// Close the database connection when the application exits
	defer config.CloseDatabase()
}
