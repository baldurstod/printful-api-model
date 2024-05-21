package schemas

type RetailCosts struct {
	Currency string  `json:"currency" bson:"currency" mapstructure:"currency"`
	Subtotal float64 `json:"subtotal" bson:"subtotal" mapstructure:"subtotal"`
	Discount float64 `json:"discount" bson:"discount" mapstructure:"discount"`
	Shipping float64 `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	Tax      float64 `json:"tax" bson:"tax" mapstructure:"tax"`
	Vat      float64 `json:"vat" bson:"vat" mapstructure:"vat"`
	Total    float64 `json:"total" bson:"total" mapstructure:"total"`
}
