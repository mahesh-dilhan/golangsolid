package main

import (
	"math"
	"testing"
)

type areatest struct {
	arg1, expected float64
}

var areatests = []areatest{
	areatest{2, 12.566371},
	areatest{3, 28.274334},
}

func TestArea(t *testing.T) {
	c := NewCircle()
	for _, test := range areatests {
		c.radius = test.arg1
		tolerance := 0.001
		output := c.Area()
		if math.Abs(output-test.expected) < tolerance {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
