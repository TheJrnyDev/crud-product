package handlers

import (
	"crud-product-bck/models"
	"crud-product-bck/services"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

// define handlers struct for product operations
type ProductHandler struct {
	productService *services.ProductService
}

// create a new product handler and return it for use
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// :GET Get all products
func (h *ProductHandler) GetProducts(c echo.Context) error {
	// Call the service to get all products
	products, err := h.productService.GetAllProducts()
	// Check for errors
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

// :POST Create a new product
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	// Bind the request body to a product model
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product data"})
	}

	// Validate the product data
	if len(product.ProductID) != 35 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID must be 35 characters long"})
	}

	// Validate the product ID format using a regular expression (XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX)
	productIdPattern := `^[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}$`
	// Compile the regular expression
	matched, err := regexp.MatchString(productIdPattern, product.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to validate product ID format"})
	}
	if !matched {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID must follow format: XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX (uppercase letters and numbers only)"})
	}

	// Check product_id is not duplicate on the database
	existingProduct, err := h.productService.GetProductByID(product.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check for existing product"})
	}
	// Product already exists, return a conflict error
	if existingProduct != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Product with this ID already exists"})
	}

	// Call the service to create a new product
	createdProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	// Return the created product with a success message
	return c.JSON(http.StatusCreated, map[string]string{"message": "Product created successfully", "product": createdProduct.ProductID})
}

// :DELETE Delete a product by ID
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	println("DeleteProduct handler called")
	// Get the product ID from the URL parameters
	productID := c.QueryParam("id")
	if productID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Product ID is required"})
	}
	println("Product ID:", productID)
	// Call the service to delete the product
	err := h.productService.DeleteProduct(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}
	// Return a success message
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
