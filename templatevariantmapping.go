package model

type TemplateVariantMapping struct {
	VariantID int64                        `json:"variant_id" bson:"variant_id"`
	Templates []TemplateVariantMappingItem `json:"templates" bson:"templates"`
}
