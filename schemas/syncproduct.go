package model

type SyncProduct struct {
	ID           int64  `json:"id" bson:"id"`
	ExternalID   string `json:"external_id" bson:"external_id"`
	Name         string `json:"name" bson:"name"`
	Variants     int    `json:"variants" bson:"variants"`
	Synced       int    `json:"synced" bson:"synced"`
	Thumbnail    string `json:"thumbnail" bson:"thumbnail"`
	ThumbnailURL string `json:"thumbnail_url" bson:"thumbnail_url"`
	IsIgnored    bool   `json:"is_ignored" bson:"is_ignored"`
}
