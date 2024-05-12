package schemas

type IncompleteItem struct {
	Name               string `json:"name" bson:"name" mapstructure:"name"`
	Quantity           int    `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	SyncVariantID      int64  `json:"sync_variant_id" bson:"sync_variant_id" mapstructure:"sync_variant_id"`
	ExternalVariantID  string `json:"external_variant_id" bson:"external_variant_id" mapstructure:"external_variant_id"`
	ExternalLineItemID string `json:"external_line_item_id" bson:"external_line_item_id" mapstructure:"external_line_item_id"`
}
