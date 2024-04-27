package model

type AddressInfo struct {
	Address1    string `json:"address1" bson:"address1" mapstructure:"address1"`
	City        string `json:"city" bson:"city" mapstructure:"city"`
	CountryCode string `json:"country_code" bson:"country_code" mapstructure:"country_code"`
	StateCode   string `json:"state_code" bson:"state_code" mapstructure:"state_code"`
	ZIP         string `json:"zip" bson:"zip" mapstructure:"zip"`
	Phone       string `json:"phone" bson:"phone" mapstructure:"phone"`
}
