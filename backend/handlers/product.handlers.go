package handlers

import (
	"crud-product-bck/models"
	"crud-product-bck/services"
	"crud-product-bck/utils"
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
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to retrieve products")
	}
	return utils.ResponseSuccess(c, http.StatusOK, products, "Products retrieved successfully")
}

// :POST Create a new product
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	// Bind the request body to a product model
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return utils.ResponseError(c, http.StatusBadRequest, "Invalid product data")
	}

	// -------------- Validate the product data section
	if len(product.ProductID) != 35 {
		return utils.ResponseError(c, http.StatusBadRequest, "Product ID must be exactly 35 characters long")
	}

	// Validate the product ID format using a regular expression (XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX)
	productIdPattern := `^[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}$`

	// Compile the regular expression
	matched, err := regexp.MatchString(productIdPattern, product.ProductID)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to validate product ID format")
	}
	if !matched {
		return utils.ResponseError(c, http.StatusBadRequest, "Product ID must be in the format XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX and contain only uppercase letters and numbers")
	}

	// Check product_id is not duplicate on the database
	existingProduct, err := h.productService.GetProductByID(product.ProductID)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to check existing product")
	}
	// Product already exists, return a conflict error
	if existingProduct != nil {
		return utils.ResponseError(c, http.StatusConflict, "Product ID already exists, please use another ID")
	}
	// ------------------ End of validation section

	// Call the service to create a new product
	createdProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to create product")
	}
	// Return the created product with a success message
	return utils.ResponseSuccess(c, http.StatusCreated, createdProduct, "Product created successfully")
}

// :DELETE Delete a product by ID
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	// Get the product ID from the URL parameters
	productID := c.QueryParam("id")
	if productID == "" {
		return utils.ResponseError(c, http.StatusBadRequest, "Query Params Product ID is required")
	}
	// Call the service to delete the product
	err := h.productService.DeleteProduct(productID)
	if err != nil {
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete product")
	}
	// Return a success message
	return utils.ResponseSuccess(c, http.StatusOK, nil, "Product deleted successfully")
}
