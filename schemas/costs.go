package schemas

type Costs struct {
	Currency          string  `json:"currency" bson:"currency" mapstructure:"currency"`
	Subtotal          float64 `json:"subtotal" bson:"subtotal" mapstructure:"subtotal"`
	Discount          float64 `json:"discount" bson:"discount" mapstructure:"discount"`
	Shipping          float64 `json:"shipping" bson:"shipping" mapstructure:"shipping"`
	Digitization      float64 `json:"digitization" bson:"digitization" mapstructure:"digitization"`
	AdditionalFee     float64 `json:"additional_fee" bson:"additional_fee" mapstructure:"additional_fee"`
	FulfillmentFee    float64 `json:"fulfillment_fee" bson:"fulfillment_fee" mapstructure:"fulfillment_fee"`
	RetailDeliveryFee string  `json:"retail_delivery_fee" bson:"retail_delivery_fee" mapstructure:"retail_delivery_fee"`
	Tax               string  `json:"tax" bson:"tax" mapstructure:"tax"`
	Vat               float64 `json:"vat" bson:"vat" mapstructure:"vat"`
	Total             float64 `json:"total" bson:"total" mapstructure:"total"`
}
