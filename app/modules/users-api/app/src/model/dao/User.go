package dao

type struct User {
	Name     string
	userId   string
	delivery []Delivery

}

type Delivery struct {
    Address  string
    Items    []Item
}