package model

type SyncVariant struct {
	ID                        int               `json:"id" bson:"id"`
	ExternalID                string            `json:"external_id" bson:"external_id"`
	SyncProductID             int               `json:"sync_product_id" bson:"sync_product_id"`
	Name                      string            `json:"name" bson:"name"`
	Synced                    bool              `json:"synced" bson:"synced"`
	VariantID                 int               `json:"variant_id" bson:"variant_id"`
	RetailPrice               string            `json:"retail_price" bson:"retail_price"`
	Currency                  string            `json:"currency" bson:"currency"`
	IsIgnored                 bool              `json:"is_ignored" bson:"is_ignored"`
	SKU                       string            `json:"sku" bson:"sku"`
	Product                   ProductVariant    `json:"product" bson:"product"`
	Files                     []SyncVariantFile `json:"files" bson:"files"`
	Options                   []ItemOption      `json:"options" bson:"options"`
	MainCategoryID            int               `json:"main_category_id" bson:"main_category_id"`
	WarehouseProductVariantID int               `json:"warehouse_product_variant_id" bson:"warehouse_product_variant_id"`
	Size                      string            `json:"size" bson:"size"`
	Color                     string            `json:"color" bson:"color"`
	AvailabilityStatus        string            `json:"availability_status" bson:"availability_status"`
}
