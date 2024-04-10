package model

type ProductVariant struct {
	VariantID int64  `json:"variant_id" bson:"variant_id"`
	ProductID int64  `json:"product_id" bson:"product_id"`
	Image     string `json:"image" bson:"image"`
	Name      string `json:"name" bson:"name"`
}
