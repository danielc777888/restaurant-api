package core

import (
	"middleearth/eateries/data"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestChangeName(t *testing.T) {
	restaurant := data.Restaurant{
		ID:   uuid.New(),
		Name: "Test",
	}
	restaurant2 := ChangeName(restaurant, "New Name")
	assert.Equal(t, restaurant2.Name, "New Name")
}
