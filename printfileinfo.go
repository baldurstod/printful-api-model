package model

type PrintfileInfo struct {
	ProductID           int                `json:"product_id" bson:"product_id"`
	AvailablePlacements interface{}        `json:"available_placements" bson:"available_placements"`
	Printfiles          []Printfile        `json:"printfiles" bson:"printfiles"`
	VariantPrintfiles   []VariantPrintfile `json:"variant_printfiles" bson:"variant_printfiles"`
	OptionGroups        []string           `json:"option_groups" bson:"option_groups"`
	Options             []string           `json:"options" bson:"options"`
}
