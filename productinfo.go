package model

type ProductInfo struct {
	Product  Product   `json:"product" bson:"product"`
	Variants []Variant `json:"variants" bson:"variants"`
}
