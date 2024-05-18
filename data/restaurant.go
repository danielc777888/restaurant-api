package data

type Restaurant struct {
	ID     uint
	Name   string `gorm:"unique"`
	Dishes []Dish
}
