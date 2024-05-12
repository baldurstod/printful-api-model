package schemas

type ItemOption struct {
	ID    string `json:"id" bson:"id"`
	Value string `json:"value" bson:"value"`
}
