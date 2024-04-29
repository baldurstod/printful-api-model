package responses

import (
	"github.com/baldurstod/printful-api-model/schemas"
)

type ShippingRates struct {
	Code   string         `json:"code" bson:"code" mapstructure:"code"`
	Result []schemas.ShippingInfo `json:"result" bson:"result" mapstructure:"result"`
}
