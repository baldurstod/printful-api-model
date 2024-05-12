package schemas

type File struct {
	Type         string       `json:"type" bson:"type" mapstructure:"type"`
	ID           int          `json:"id" bson:"id" mapstructure:"id"`
	URL          string       `json:"url" bson:"url" mapstructure:"url"`
	Options      []FileOption `json:"options" bson:"options" mapstructure:"options"`
	Hash         string       `json:"hash" bson:"hash" mapstructure:"hash"`
	Filename     string       `json:"filename" bson:"filename" mapstructure:"filename"`
	MimeType     string       `json:"mime_type" bson:"mime_type" mapstructure:"mime_type"`
	Size         int          `json:"size" bson:"size" mapstructure:"size"`
	Width        int          `json:"width" bson:"width" mapstructure:"width"`
	Height       int          `json:"height" bson:"height" mapstructure:"height"`
	DPI          int          `json:"dpi" bson:"dpi" mapstructure:"dpi"`
	Status       string       `json:"status" bson:"status" mapstructure:"status"`
	Created      int64        `json:"created" bson:"created" mapstructure:"created"`
	ThumbnailURL string       `json:"thumbnail_url" bson:"thumbnail_url" mapstructure:"thumbnail_url"`
	PreviewURL   string       `json:"preview_url" bson:"preview_url" mapstructure:"preview_url"`
	Visible      bool         `json:"visible" bson:"visible" mapstructure:"visible"`
	IsTemporary  bool         `json:"is_temporary" bson:"is_temporary" mapstructure:"is_temporary"`
}
