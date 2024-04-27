package model

type CalculateShippingRatesRequest struct {
	Recipient AddressInfo `json:"recipient" bson:"recipient" mapstructure:"recipient"`
	Items     []ItemInfo  `json:"items" bson:"items" mapstructure:"items"`
	Currency  string      `json:"currency" bson:"currency" mapstructure:"currency"`
	Locale    string      `json:"locale" bson:"locale" mapstructure:"locale"`
}
