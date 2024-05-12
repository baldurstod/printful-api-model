package schemas

type Costs struct {
	Currency          string `json:"currency" bson:"currency" mapstructure:"currency"`
	Subtotal          string `json:"subtotal" bson:"subtotal" mapstructure:"subtotal"`
	Discount          string `json:"discount" bson:"discount" mapstructure:"discount"`
	Shipping          string `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	Digitization      string `json:"digitization" bson:"digitization" mapstructure:"digitization"`
	AdditionalFee     string `json:"additional_fee" bson:"additional_fee" mapstructure:"additional_fee"`
	FulfillmentFee    string `json:"fulfillment_fee" bson:"fulfillment_fee" mapstructure:"fulfillment_fee"`
	RetailDeliveryFee string `json:"retail_delivery_fee" bson:"retail_delivery_fee" mapstructure:"retail_delivery_fee"`
	Tax               string `json:"tax" bson:"tax" mapstructure:"tax"`
	Vat               string `json:"vat" bson:"vat" mapstructure:"vat"`
	Total             string `json:"total" bson:"total" mapstructure:"total"`
}
