package handlers

import (
	"crud-product-bck/messages"
	"crud-product-bck/models"
	"crud-product-bck/services"
	"crud-product-bck/utils"
	"log"
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
		log.Println("Error to get all product data: ", err)
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to retrieve products")
	}

	return utils.ResponseSuccess(c, http.StatusOK, products, "Products retrieved successfully")
}

// :POST Create a new product
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	// Bind the request body to a product model
	var product models.Product
	if err := c.Bind(&product); err != nil {
		log.Println("Error to bind data for create new product: ", err)
		return utils.ResponseError(c, http.StatusBadRequest, messages.BadRequest)
	}

	// -------------- Validate the product data section
	if len(product.ProductID) != 35 {
		log.Println("Invalid productID, must be 35 characters long")
		return utils.ResponseError(c, http.StatusBadRequest, "Product ID must be exactly 35 characters long")
	}

	// Validate the product ID format using a regular expression (XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX)
	productIdPattern := `^[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}$`

	// Compile the regular expression
	matched, err := regexp.MatchString(productIdPattern, product.ProductID)
	if err != nil {
		log.Println("Error to validate productID: ", err)
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to validate product ID format")
	}
	if !matched {
		log.Println("Invalid productID format")
		return utils.ResponseError(c, http.StatusBadRequest, "Product ID must be in the format XXXXX-XXXXX-XXXXX-XXXXX-XXXXX-XXXXX and contain only uppercase letters and numbers")
	}

	// Check product_id is not duplicate on the database
	existingProduct, err := h.productService.GetProductByID(product.ProductID)
	if err != nil {
		log.Println("Failed to get product by productID: ", err)
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to check existing product")
	}
	// Product already exists, return a conflict error
	if existingProduct != nil {
		log.Println("Failed to create product, This productID already exists")
		return utils.ResponseError(c, http.StatusConflict, "Product ID already exists, please use another ID")
	}
	// ------------------ End of validation section

	// Call the service to create a new product
	createdProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		log.Println("Failed to create product on database: ", err)
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
		log.Println("Invalid query params")
		return utils.ResponseError(c, http.StatusBadRequest, "Query Params Product ID is required")
	}
	// Call the service to delete the product
	err := h.productService.DeleteProduct(productID)
	if err != nil {
		log.Println("Failed to delete product on database: ", err)
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to delete product")
	}
	// Return a success message
	return utils.ResponseSuccess(c, http.StatusOK, nil, "Product deleted successfully")
}

// :PUT Update product name by productID
func (h *ProductHandler) UpdateProductName(c echo.Context) error {
	// Get the product ID from the URL paramters
	productID := c.QueryParam("id")
	if productID == "" {
		log.Println("Invalid query params")
		return utils.ResponseError(c, http.StatusBadRequest, "Query Params Product ID is required")
	}

	// Bind the request
	var product models.Product
	err := c.Bind(&product)
	if err != nil {
		log.Println("[product-update] Invalid body request")
		return utils.ResponseError(c, http.StatusBadRequest, messages.BadRequest)
	}

	// Call update product name service
	updateData, err := h.productService.UpdateProductName(productID, product.ProductName)
	if err != nil {
		log.Println("Error to update product: ", err)
		return utils.ResponseError(c, http.StatusInternalServerError, "Failed to update product")
	}

	// Check data has update or not?
	if updateData.MatchedCount <= 0 {
		// return response error to client
		log.Println("Not found product to update")
		return utils.ResponseError(c, http.StatusNotFound, "Not found product to update")
	}

	return utils.ResponseSuccess(c, http.StatusOK, updateData, "update successfully")
}
