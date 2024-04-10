package model

type Variant struct {
	ID                  int64                `json:"id" bson:"id"`
	ProductID           int64                `json:"product_id" bson:"product_id"`
	Name                string               `json:"name" bson:"name"`
	Size                string               `json:"size" bson:"size"`
	Color               string               `json:"color" bson:"color"`
	ColorCode           string               `json:"color_code" bson:"color_code"`
	ColorCode2          string               `json:"color_code2" bson:"color_code2"`
	Image               string               `json:"image" bson:"image"`
	Price               string               `json:"price" bson:"price"`
	inStock             bool                 `json:"in_stock" bson:"in_stock"`
	AvailabilityRegions map[string]string    `json:"availability_regions" bson:"availability_regions"`
	AvailabilityStatus  []AvailabilityStatus `json:"availability_status" bson:"availability_status"`
	Material            []Material           `json:"material" bson:"material"`
}
