package model

type SyncProductInfo struct {
	SyncProduct  SyncProduct   `json:"sync_product" bson:"sync_product"`
	SyncVariants []SyncVariant `json:"sync_variants" bson:"sync_variants"`
}
