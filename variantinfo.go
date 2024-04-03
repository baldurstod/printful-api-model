package model

type VariantInfo struct {
	Variant `json:"variant" bson:"variant"`
	Product `json:"product" bson:"product"`
}
