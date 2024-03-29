package model

type OptionType struct {
	ID                       string            `json:"id"`
	Title                    string            `json:"title"`
	Type                     string            `json:"type"`
	Values                   map[string]string `json:"values"`
	AdditionalPrice          string            `json:"additional_price"`
	AdditionalPriceBreakdown map[string]string `json:"additional_price_breakdown"`
}
