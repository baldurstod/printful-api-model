package model

type Printfile struct {
	PrintfileID int64  `json:"printfile_id" bson:"printfile_id"`
	Width       int    `json:"width" bson:"width"`
	Height      int    `json:"height" bson:"height"`
	DPI         int    `json:"dpi" bson:"dpi"`
	FillMode    string `json:"fill_mode" bson:"fill_mode"`
	CanRotate   bool   `json:"can_rotate" bson:"can_rotate"`
}
