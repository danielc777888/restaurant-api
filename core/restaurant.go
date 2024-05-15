package core

import "middleearth/eateries/data"

func ChangeName(restaurant data.Restaurant, name string) data.Restaurant {
	restaurant.Name = name
	return restaurant
}
