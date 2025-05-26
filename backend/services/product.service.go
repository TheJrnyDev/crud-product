package services

import (
	"context"
	"crud-product-bck/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	collection *mongo.Collection
}

func NewProductService(client *mongo.Client, dbName string) *ProductService {
	// define collection name
	collectionName := "products"

	return &ProductService{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

// GetAllProducts retrieves all products from the collection
func (s *ProductService) GetAllProducts() (*[]models.Product, error) {
	// Find all documents in the collection
	cursor, err := s.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	// Ensure the cursor is closed after use
	defer cursor.Close(context.TODO())

	// Get all results from the cursor and unmarshal into a slice of Product models
	var results []models.Product
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// GetProductByID retrieves a product by its ID from the collection
func (s *ProductService) GetProductByID(productID string) (*models.Product, error) {
	// Find a single document with the specified product ID
	var product models.Product
	err := s.collection.FindOne(context.TODO(), bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No product found with the given ID
		}
		return nil, err // Return any other error
	}

	return &product, nil
}

// CreateProduct inserts a new product into the collection
func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	// Insert the product into the collection
	_, err := s.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}

	// Return the created product
	return product, nil
}

// DeleteProduct removes a product from the collection by its ID
func (s *ProductService) DeleteProduct(productID string) error {
	// Delete the product with the specified ID
	_, err := s.collection.DeleteOne(context.TODO(), bson.M{"product_id": productID})
	if err != nil {
		return err
	}

	return nil
}
