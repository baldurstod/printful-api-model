package model

type ProductVariant struct {
	VariantID int    `json:"variant_id" bson:"variant_id"`
	ProductID int    `json:"product_id" bson:"product_id"`
	Image     string `json:"image" bson:"image"`
	Name      string `json:"name" bson:"name"`
}
