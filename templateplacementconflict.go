package model

type TemplatePlacementConflict struct {
	Placement string   `json:"placement" bson:"placement"`
	Conflicts []string `json:"conflicts" bson:"conflicts"`
}
