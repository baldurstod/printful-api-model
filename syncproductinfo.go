package model

import (
	"github.com/baldurstod/printful-api-model/schemas"
)

type SyncProductInfo struct {
	SyncProduct  schemas.SyncProduct   `json:"sync_product" bson:"sync_product"`
	SyncVariants []schemas.SyncVariant `json:"sync_variants" bson:"sync_variants"`
}
