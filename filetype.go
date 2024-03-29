package model

type FileType struct {
	ID              string              `json:"id"`
	Type            string              `json:"type"`
	Title           string              `json:"title"`
	AdditionalPrice string              `json:"additional_price"`
	Options         []CatalogFileOption `json:"options"`
}
