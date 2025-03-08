package dao

type User struct {
	name         string
	userId       string
	deliveryList []Delivery
}

type Delivery struct {
	deliveryId        string
	initialCoordinate Coordinate
	finalCoordinate   Coordinate
}

type Coordinate struct {
	latitude  float32
	longitude float32
}
