package data

type Dish struct {
	ID           uint
	Name         string
	Description  string
	Price        uint
	Ratings      []Rating
	RestaurantID uint
}
