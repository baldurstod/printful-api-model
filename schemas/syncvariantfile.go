package model

type SyncVariantFile struct {
	Type            string       `json:"type" bson:"type"`
	ID              int          `json:"id" bson:"id"`
	URL             string       `json:"url" bson:"url"`
	Options         []FileOption `json:"options" bson:"options"`
	Hash            string       `json:"hash" bson:"hash"`
	Filename        string       `json:"filename" bson:"filename"`
	MimeType        string       `json:"mime_type" bson:"mime_type"`
	Size            int          `json:"size" bson:"size"`
	Width           int          `json:"width" bson:"width"`
	Height          int          `json:"height" bson:"height"`
	DPI             int          `json:"dpi" bson:"dpi"`
	Status          string       `json:"status" bson:"status"`
	Created         int64        `json:"created" bson:"created"`
	ThumbnailURL    string       `json:"thumbnail_url" bson:"thumbnail_url"`
	PreviewURL      string       `json:"preview_url" bson:"preview_url"`
	Visible         bool         `json:"visible" bson:"visible"`
	IsTemporary     bool         `json:"is_temporary" bson:"is_temporary"`
	StitchCountTier string       `json:"stitch_count_tier" bson:"stitch_count_tier"`
}
