package model

type VariantPrintfile struct {
	VariantID  int64       `json:"variant_id" bson:"variant_id"`
	Placements interface{} `json:"placements" bson:"placements"`
}
