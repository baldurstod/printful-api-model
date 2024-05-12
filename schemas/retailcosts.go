package schemas

type RetailCosts struct {
	Currency string `json:"currency" bson:"currency" mapstructure:"currency"`
	Subtotal string `json:"subtotal" bson:"subtotal" mapstructure:"subtotal"`
	Discount string `json:"discount" bson:"discount" mapstructure:"discount"`
	Shipping string `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	Tax      string `json:"tax" bson:"tax" mapstructure:"tax"`
	Vat      string `json:"vat" bson:"vat" mapstructure:"vat"`
	Total    string `json:"total" bson:"total" mapstructure:"total"`
}
