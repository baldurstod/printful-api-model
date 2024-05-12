package schemas

type ProductVariant struct {
	VariantID int    `json:"variant_id" bson:"variant_id" mapstructure:"variant_id"`
	ProductID int    `json:"product_id" bson:"product_id" mapstructure:"product_id"`
	Image     string `json:"image" bson:"image" mapstructure:"image"`
	Name      string `json:"name" bson:"name" mapstructure:"name"`
}
