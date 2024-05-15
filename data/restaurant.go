package data

type Restaurant struct {
	ID     uint
	Name   string
	Dishes []Dish
}

type Dish struct {
	ID           uint
	Name         string
	Description  string
	Price        uint
	RestaurantID uint
}
