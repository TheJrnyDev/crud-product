package models

type Product struct {
	ProductID   string `json:"product_id" bson:"product_id"`
	ProductName string `json:"product_name bson:"product_name"`
}
