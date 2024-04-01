package model

type VariantResponse struct {
	Success bool `json:"success"`
	Variant `json:"variant"`
}
