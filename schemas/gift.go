package schemas

type Gift struct {
	Subject string `json:"subject" bson:"subject" mapstructure:"subject"`
	Message string `json:"message" bson:"message" mapstructure:"message"`
}
