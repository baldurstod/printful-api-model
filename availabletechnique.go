package model

type AvailableTechnique struct {
	Key         string `json:"key"`
	DisplayName string `json:"display_name"`
	IsDefault   bool   `json:"is_default"`
}
