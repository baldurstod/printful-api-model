package model

type ProductResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Product  Product   `json:"product"`
		Variants []Variant `json:"variants"`
	} `json:"result"`
}
