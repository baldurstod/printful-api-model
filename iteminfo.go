package model

type ItemInfo struct {
	VariantID                 string `json:"variant_id" bson:"variant_id" mapstructure:"variant_id"`
	ExternalVariantID         string `json:"external_variant_id" bson:"external_variant_id" mapstructure:"external_variant_id"`
	WarehouseProductVariantID string `json:"warehouse_product_variant_id" bson:"warehouse_product_variant_id" mapstructure:"warehouse_product_variant_id"`
	Quantity                  int    `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	Value                     string `json:"value" bson:"value" mapstructure:"value"`
}
