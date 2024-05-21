package schemas

type Item struct {
	ID                        int64          `json:"id,omitempty" bson:"id" mapstructure:"id"`
	ExternalID                string         `json:"external_id,omitempty" bson:"external_id" mapstructure:"external_id"`
	VariantID                 int64          `json:"variant_id,omitempty" bson:"variant_id" mapstructure:"variant_id"`
	SyncVariantID             int64          `json:"sync_variant_id,omitempty" bson:"sync_variant_id" mapstructure:"sync_variant_id"`
	ExternalVariantID         string         `json:"external_variant_id,omitempty" bson:"external_variant_id" mapstructure:"external_variant_id"`
	WarehouseProductVariantID int64          `json:"warehouse_product_variant_id,omitempty" bson:"warehouse_product_variant_id" mapstructure:"warehouse_product_variant_id"`
	ProductTemplateID         int64          `json:"product_template_id,omitempty" bson:"product_template_id" mapstructure:"product_template_id"`
	ExternalProductID         string         `json:"external_product_id,omitempty" bson:"external_product_id" mapstructure:"external_product_id"`
	Quantity                  int            `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	Price                     string         `json:"price" bson:"price" mapstructure:"price"`
	RetailPrice               string         `json:"retail_price" bson:"retail_price" mapstructure:"retail_price"`
	Name                      string         `json:"name" bson:"name" mapstructure:"name"`
	Product                   ProductVariant `json:"product" bson:"product" mapstructure:"product"`
	Files                     []File         `json:"files,omitempty" bson:"files" mapstructure:"files"`
	Options                   []Option       `json:"options,omitempty" bson:"options" mapstructure:"options"`
	SKU                       string         `json:"sku" bson:"sku" mapstructure:"sku"`
	Discontinued              bool           `json:"discontinued" bson:"discontinued" mapstructure:"discontinued"`
	OutOfStock                bool           `json:"out_of_stock" bson:"out_of_stock" mapstructure:"out_of_stock"`
}

func NewItem() Item {
	return Item{
		Files:   make([]File, 0),
		Options: make([]Option, 0),
	}
}
