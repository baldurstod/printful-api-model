package model

type CatalogFileOption struct {
	ID              string  `json:"id"`
	Type            string  `json:"type"`
	Title           string  `json:"title"`
	AdditionalPrice float64 `json:"additional_price"`
}
