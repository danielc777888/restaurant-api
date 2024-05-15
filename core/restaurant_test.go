package core

import (
	"middleearth/eateries/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeName(t *testing.T) {
	restaurant := data.Restaurant{
		ID:   1,
		Name: "Test",
	}
	restaurant2 := ChangeName(restaurant, "New Name")
	assert.Equal(t, restaurant2.Name, "New Names")
}
