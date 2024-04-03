package model

type VariantPrintfile struct {
	VariantID  int         `json:"variant_id" bson:"variant_id"`
	Placements interface{} `json:"placements" bson:"placements"`
}
