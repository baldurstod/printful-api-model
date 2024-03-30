package model

type OptionType struct {
	ID                       string      `json:"id"`
	Title                    string      `json:"title"`
	Type                     string      `json:"type"`
	Values                   interface{} `json:"values"`
	AdditionalPrice          string      `json:"additional_price"`
	AdditionalPriceBreakdown interface{} `json:"additional_price_breakdown"`
}
