package utils

import "fmt"

type Point2d struct {
	X int
	Y int
}

func (p Point2d) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
