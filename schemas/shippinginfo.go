package schemas

type ShippingInfo struct {
	ID              string `json:"id" bson:"id" mapstructure:"id"`
	Name            string `json:"name" bson:"name" mapstructure:"name"`
	Rate            string `json:"rate" bson:"rate" mapstructure:"rate"`
	Currency        string `json:"currency" bson:"currency" mapstructure:"currency"`
	MinDeliveryDays int    `json:"minDeliveryDays" bson:"minDeliveryDays" mapstructure:"minDeliveryDays"`
	MaxDeliveryDays int    `json:"maxDeliveryDays" bson:"maxDeliveryDays" mapstructure:"maxDeliveryDays"`
	MinDeliveryDate string `json:"minDeliveryDate" bson:"minDeliveryDate" mapstructure:"minDeliveryDate"`
	MaxDeliveryDate string `json:"maxDeliveryDate" bson:"maxDeliveryDate" mapstructure:"maxDeliveryDate"`
}
