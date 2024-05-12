package schemas

type FileOption struct {
	ID    string `json:"id" bson:"id" mapstructure:"id"`
	Value string `json:"value" bson:"value" mapstructure:"value"`
}
