package model

type Template struct {
	TemplateID        int    `json:"template_id" bson:"template_id"`
	ImageURL          string `json:"image_url" bson:"image_url"`
	BackgroundURL     string `json:"background_url" bson:"background_url"`
	BackgroundColor   string `json:"background_color" bson:"background_color"`
	PrintfileID       int    `json:"printfile_id" bson:"printfile_id"`
	TemplateWidth     int    `json:"template_width" bson:"template_width"`
	TemplateHeight    int    `json:"template_height" bson:"template_height"`
	PrintAreaWidth    int    `json:"print_area_width" bson:"print_area_width"`
	PrintAreaHeight   int    `json:"print_area_height" bson:"print_area_height"`
	PrintAreaTop      int    `json:"print_area_top" bson:"print_area_top"`
	PrintAreaLeft     int    `json:"print_area_left" bson:"print_area_left"`
	IsTemplateOnFront bool   `json:"is_template_on_front" bson:"is_template_on_front"`
	Orientation       string `json:"orientation" bson:"orientation"`
}
