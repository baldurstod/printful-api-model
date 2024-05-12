package schemas

type PackingSlip struct {
	Email         string `json:"email" bson:"email" mapstructure:"email"`
	Phone         string `json:"phone" bson:"phone" mapstructure:"phone"`
	Message       string `json:"message" bson:"message" mapstructure:"message"`
	LogoURL       string `json:"logo_url" bson:"logo_url" mapstructure:"logo_url"`
	StoreName     string `json:"store_name" bson:"store_name" mapstructure:"store_name"`
	CustomOrderID string `json:"custom_order_id" bson:"custom_order_id" mapstructure:"custom_order_id"`
}
