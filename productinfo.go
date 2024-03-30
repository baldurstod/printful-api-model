package model

type ProductInfo struct {
	Product  `json:"product"`
	Variants []Variant `json:"variants"`
}
