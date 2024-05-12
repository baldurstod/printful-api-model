package schemas

type OrderShipmentItem struct {
	ItemID   int `json:"item_id" bson:"item_id" mapstructure:"item_id"`
	Quantity int `json:"quantity" bson:"quantity" mapstructure:"quantity"`
	Picked   int `json:"picked" bson:"picked" mapstructure:"picked"`
	Printed  int `json:"printed" bson:"printed" mapstructure:"printed"`
}
