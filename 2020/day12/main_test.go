package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveNorth(t *testing.T) {
	boat := ferry{x: 0, y: 0, direction: east}
	boat.turn(-90)
	assert.Equal(t, boat.direction, north)
	boat.turn(-270)
	assert.Equal(t, boat.direction, east)
	boat.turn(270)
	assert.Equal(t, boat.direction, north)
	boat.turn(270)
	assert.Equal(t, boat.direction, west)
	boat.turn(180)
	assert.Equal(t, boat.direction, east)
	boat.turn(90)
	assert.Equal(t, boat.direction, south)
	boat.turn(-180)
	assert.Equal(t, boat.direction, north)
}

func TestAngleStep2(t *testing.T) {
	x, y := wayPointTurn(4, 10, 90)
	assert.Equal(t, 10, x)
	assert.Equal(t, -4, y)

	x, y = wayPointTurn(4, 10, 180)
	assert.Equal(t, -4, x)
	assert.Equal(t, -10, y)

	x, y = wayPointTurn(4, 10, 270)
	assert.Equal(t, -10, x)
	assert.Equal(t, 4, y)

	x, y = wayPointTurn(4, 10, -90)
	assert.Equal(t, -10, x)
	assert.Equal(t, 4, y)

	x, y = wayPointTurn(4, 10, -180)
	assert.Equal(t, -4, x)
	assert.Equal(t, -10, y)

	x, y = wayPointTurn(4, 10, -270)
	assert.Equal(t, 10, x)
	assert.Equal(t, -4, y)

}
