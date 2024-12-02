package models



type ShippingInfo struct {
	Address     string `bson:"address" json:"address"`         // Shipping Address
	City        string `bson:"city" json:"city"`               // City
	State       string `bson:"state" json:"state"`             // State
	PostalCode  string `bson:"postal_code" json:"postal_code"` // Postal Code
	Country     string `bson:"country" json:"country"`         // Country
	PhoneNumber string `bson:"phone_number" json:"phone_number"` // Contact Number
}
