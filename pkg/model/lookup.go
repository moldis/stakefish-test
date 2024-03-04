package model

type Address struct {
	IP string `json:"ip" bson:"ip"`
}

type Query struct {
	Addresses []Address `json:"addresses" bson:"addresses"`
	ClientIP  string    `json:"client_ip" bson:"client_ip"`
	CreatedAt int64     `json:"created_at" bson:"created_at"`
	Domain    string    `json:"domain" bson:"domain"`
}
