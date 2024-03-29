package model

type ProductResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Product `json:"pf"`
	} `json:"result"`
}

type Product struct {
	ID                 int                  `json:"id"`
	MainCategoryID     int                  `json:"main_category_id"`
	Type               string               `json:"type"`
	TypeName           string               `json:"type_name"`
	Title              string               `json:"title"`
	Brand              string               `json:"brand"`
	Model              string               `json:"model"`
	Image              string               `json:"image"`
	VariantCount       int                  `json:"variant_count"`
	Currency           string               `json:"currency"`
	Files              []FileType           `json:"files"`
	Options            []OptionType         `json:"options"`
	IsDiscontinued     bool                 `json:"is_discontinued"`
	AvgFulfillmentTime float64              `json:"avg_fulfillment_time"`
	Description        string               `json:"description"`
	Techniques         []AvailableTechnique `json:"techniques"`
	OriginCountry      string               `json:"origin_country"`
}
