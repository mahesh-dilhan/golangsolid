package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) Area() {
	fmt.Printf("Area of: %f\n ", math.Pi*c.radius*c.radius)
}

func NewCircle() circle {
	c := &circle{}
	return *c
}

func main() {
	c := NewCircle()
	c.radius = 3.0
	c.Area()
}
