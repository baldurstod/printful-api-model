package model

type ProductTemplate struct {
	Version               int                         `json:"version" bson:"version"`
	MinDPI                int                         `json:"min_dpi" bson:"min_dpi"`
	VariantMapping        []TemplateVariantMapping    `json:"variant_mapping" bson:"variant_mapping"`
	Templates             []Template                  `json:"templates" bson:"templates"`
	ConflictingPlacements []TemplatePlacementConflict `json:"conflicting_placements" bson:"conflicting_placements"`
}
