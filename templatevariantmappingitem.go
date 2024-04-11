package model

type TemplateVariantMappingItem struct {
	Placement  string `json:"placement" bson:"placement"`
	TemplateID int    `json:"template_id" bson:"template_id"`
}
