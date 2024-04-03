package model

type VariantResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Variant `json:"variant"`
		Product `json:"product"`
	} `json:"result"`
}
