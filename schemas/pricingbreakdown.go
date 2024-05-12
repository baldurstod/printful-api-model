package schemas

type PricingBreakdown struct {
	CustomerPays   string `json:"customer_pays" bson:"customer_pays" mapstructure:"customer_pays"`
	PrintfulPrice  string `json:"printful_price" bson:"printful_price" mapstructure:"printful_price"`
	Profit         string `json:"profit" bson:"profit" mapstructure:"profit"`
	CurrencySymbol string `json:"currency_symbol" bson:"currency_symbol" mapstructure:"currency_symbol"`
}
