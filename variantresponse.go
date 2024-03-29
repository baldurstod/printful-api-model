package model

type VariantResponse struct {
	Success bool `json:"success"`
	Variant `json:"variant"`
}

type Variant struct {
	ID                  int                  `json:"id"`
	ProductID           int                  `json:"product_id"`
	Name                string               `json:"name"`
	Size                string               `json:"size"`
	Color               string               `json:"color"`
	ColorCode           string               `json:"color_code"`
	ColorCode2          string               `json:"color_code2"`
	Image               string               `json:"image"`
	Price               string               `json:"price"`
	inStock             bool                 `json:"in_stock"`
	AvailabilityRegions map[string]string    `json:"availability_regions"`
	AvailabilityStatus  []AvailabilityStatus `json:"availability_status"`
	Material            []Material           `json:"material"`
}
