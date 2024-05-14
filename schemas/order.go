package schemas

type Order struct {
	ID                  int64              `json:"id" bson:"id" mapstructure:"id"`
	ExternalID          string             `json:"external_id" bson:"external_id" mapstructure:"external_id"`
	Store               int64              `json:"store" bson:"store" mapstructure:"store"`
	Status              string             `json:"status" bson:"status" mapstructure:"status"`
	Shipping            string             `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	ShippingServiceName string             `json:"shipping_service_name" bson:"shipping_service_name" mapstructure:"shipping_service_name"`
	Created             int64              `json:"created" bson:"created" mapstructure:"created"`
	Updated             int64              `json:"updated" bson:"updated" mapstructure:"updated"`
	Recipient           Address            `json:"recipient" bson:"recipient" mapstructure:"recipient"`
	Items               []Item             `json:"items" bson:"items" mapstructure:"items"`
	BrandingItems       []Item             `json:"branding_items" bson:"branding_items" mapstructure:"branding_items"`
	IncompleteItems     []IncompleteItem   `json:"incomplete_items" bson:"incomplete_items" mapstructure:"incomplete_items"`
	Costs               Costs              `json:"costs" bson:"costs" mapstructure:"costs"`
	RetailCosts         RetailCosts        `json:"retail_costs" bson:"retail_costs" mapstructure:"retail_costs"`
	PricingBreakdown    []PricingBreakdown `json:"pricing_breakdown" bson:"pricing_breakdown" mapstructure:"pricing_breakdown"`
	Shipment            []Shipment         `json:"shipments" bson:"shipments" mapstructure:"shipments"`
	Gift                Gift               `json:"gift" bson:"gift" mapstructure:"gift"`
	PackingSlip         PackingSlip        `json:"packing_slip" bson:"packing_slip" mapstructure:"packing_slip"`
}

func NewOrder() Order {
	return Order{
		Items:            make([]Item, 0),
		BrandingItems:    make([]Item, 0),
		IncompleteItems:  make([]IncompleteItem, 0),
		PricingBreakdown: make([]PricingBreakdown, 0),
		Shipment:         make([]Shipment, 0),
	}
}
