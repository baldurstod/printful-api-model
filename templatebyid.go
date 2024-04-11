package model

type TemplateById struct {
	ID                  int           `json:"id" bson:"id"`
	ProductID           int           `json:"product_id" bson:"product_id"`
	ExternalProductID   int           `json:"external_product_id" bson:"external_product_id"`
	Title               string        `json:"title" bson:"title"`
	AvailableVariantIDs []int         `json:"available_variant_ids" bson:"available_variant_ids"`
	OptionData          interface{}   `json:"option_data" bson:"option_data"`
	Colors              interface{}   `json:"colors" bson:"colors"`
	Sizes               []string      `json:"sizes" bson:"sizes"`
	MockupFileURL       string        `json:"mockup_file_url" bson:"mockup_file_url"`
	Placements          []interface{} `json:"placements" bson:"placements"`
	CreatedAt           int           `json:"created_at" bson:"created_at"`
	UpdatedAt           int           `json:"updated_at" bson:"updated_at"`
	PlacementOptionData []interface{} `json:"placement_option_data" bson:"placement_option_data"`
}
