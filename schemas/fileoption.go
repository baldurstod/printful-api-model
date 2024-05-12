package model

type FileOption struct {
	ID    string `json:"id" bson:"id"`
	Value string `json:"value" bson:"value"`
}
