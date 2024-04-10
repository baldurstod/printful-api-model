package model

type Product struct {
	ID                 int64                `json:"id" bson:"id"`
	MainCategoryID     int                  `json:"main_category_id" bson:"main_category_id"`
	Type               string               `json:"type" bson:"type"`
	TypeName           string               `json:"type_name" bson:"type_name"`
	Title              string               `json:"title" bson:"title"`
	Brand              string               `json:"brand" bson:"brand"`
	Model              string               `json:"model" bson:"model"`
	Image              string               `json:"image" bson:"image"`
	VariantCount       int                  `json:"variant_count" bson:"variant_count"`
	Currency           string               `json:"currency" bson:"currency"`
	Files              []FileType           `json:"files" bson:"files"`
	Options            []OptionType         `json:"options" bson:"options"`
	IsDiscontinued     bool                 `json:"is_discontinued" bson:"is_discontinued"`
	AvgFulfillmentTime float64              `json:"avg_fulfillment_time" bson:"avg_fulfillment_time"`
	Description        string               `json:"description" bson:"description"`
	Techniques         []AvailableTechnique `json:"techniques" bson:"techniques"`
	OriginCountry      string               `json:"origin_country" bson:"origin_country"`
}
