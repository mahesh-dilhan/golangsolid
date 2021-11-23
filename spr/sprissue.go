package spr

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() {
	fmt.Printf("Area of: %f\n ", math.Pi*c.radius*c.radius)
}
