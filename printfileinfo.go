package model

type PrintfileInfo struct {
	ProductID           int                `json:"product_id" bson:"product_id"`
	AvailablePlacements interface{}        `json:"available_placements" bson:"available_placements"`
	Printfiles          []Printfile        `json:"printfiles" bson:"printfiles"`
	VariantPrintfiles   []VariantPrintfile `json:"variant_printfiles" bson:"variant_printfiles"`
	OptionGroups        []string           `json:"option_groups" bson:"option_groups"`
	Options             []string           `json:"options" bson:"options"`
}

func (printfileInfo PrintfileInfo) GetPrintfile(variantID int, placement string) *Printfile {
	for _, v := range printfileInfo.VariantPrintfiles {
		if v.VariantID == variantID {

			printfileID, ok := v.Placements.(map[string]interface{})[placement]
			if !ok {
				return nil
			}

			for _, p := range printfileInfo.Printfiles {
				if p.PrintfileID == int(printfileID.(float64)) {
					return &p
				}
			}
			return nil
		}
	}
	return nil
}
