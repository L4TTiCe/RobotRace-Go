package utils

import (
	"errors"
	"fmt"
)

type Direction struct {
	name               string
	horizontalModifier int
	verticalModifier   int
}

func (direction *Direction) String() string {
	return fmt.Sprintf("%s", direction.name)
}

var NORTH *Direction
var SOUTH *Direction
var EAST *Direction
var WEST *Direction

func init() {
	NORTH = &Direction{
		name:               "NORTH",
		horizontalModifier: 0,
		verticalModifier:   1,
	}
	SOUTH = &Direction{
		name:               "SOUTH",
		horizontalModifier: 0,
		verticalModifier:   -1,
	}
	EAST = &Direction{
		name:               "EAST",
		horizontalModifier: 1,
		verticalModifier:   0,
	}
	WEST = &Direction{
		name:               "WEST",
		horizontalModifier: -1,
		verticalModifier:   0,
	}
}

func GetNorth() *Direction {
	return NORTH
}

func GetSouth() *Direction {
	return SOUTH
}

func GetEast() *Direction {
	return EAST
}

func GetWest() *Direction {
	return WEST
}

func (d *Direction) GetHorizontalModifier() int {
	return d.horizontalModifier
}

func (d *Direction) GetVerticalModifier() int {
	return d.verticalModifier
}

func (d *Direction) Complement() (*Direction, error) {
	switch d {
	case NORTH:
		return SOUTH, nil
	case SOUTH:
		return NORTH, nil
	case EAST:
		return WEST, nil
	case WEST:
		return EAST, nil
	default:
		return nil, errors.New("illegal Direction")
	}
}

func (d *Direction) GetName() string {
	return d.name
}

func (d *Direction) GetRight() (*Direction, error) {
	switch d {
	case NORTH:
		return EAST, nil
	case EAST:
		return SOUTH, nil
	case SOUTH:
		return WEST, nil
	case WEST:
		return NORTH, nil
	default:
		return nil, errors.New("illegal Direction")
	}
}

func (d *Direction) GetLeft() (*Direction, error) {
	switch d {
	case NORTH:
		return WEST, nil
	case WEST:
		return SOUTH, nil
	case SOUTH:
		return EAST, nil
	case EAST:
		return SOUTH, nil
	default:
		return nil, errors.New("illegal Direction")
	}
}
