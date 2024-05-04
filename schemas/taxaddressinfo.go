package schemas

type TaxAddressInfo struct {
	City        string `json:"city" bson:"city" mapstructure:"city"`
	CountryCode string `json:"country_code" bson:"country_code" mapstructure:"country_code"`
	StateCode   string `json:"state_code" bson:"state_code" mapstructure:"state_code"`
	ZIP         string `json:"zip" bson:"zip" mapstructure:"zip"`
}
