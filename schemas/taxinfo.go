package schemas

type TaxInfo struct {
	Required        bool    `json:"required" bson:"required" mapstructure:"required"`
	Rate            float64 `json:"rate" bson:"rate" mapstructure:"rate"`
	ShippingTaxable bool    `json:"shipping_taxable" bson:"shipping_taxable" mapstructure:"shipping_taxable"`
}
