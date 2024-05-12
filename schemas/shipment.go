package schemas

type Shipment struct {
	ID             int64               `json:"id" bson:"id" mapstructure:"id"`
	Carrier        string              `json:"carrier" bson:"carrier" mapstructure:"carrier"`
	Service        string              `json:"service" bson:"service" mapstructure:"service"`
	TrackingNumber string              `json:"tracking_number" bson:"tracking_number" mapstructure:"tracking_number"`
	TrackingURL    string              `json:"tracking_url" bson:"tracking_url" mapstructure:"tracking_url"`
	Created        int64               `json:"created" bson:"created" mapstructure:"created"`
	ShipDate       string              `json:"ship_date" bson:"ship_date" mapstructure:"ship_date"`
	ShippedAt      string              `json:"shipped_at" bson:"shipped_at" mapstructure:"shipped_at"`
	Reshipment     bool                `json:"reshipment" bson:"reshipment" mapstructure:"reshipment"`
	items          []OrderShipmentItem `json:"items" bson:"items" mapstructure:"items"`
}
