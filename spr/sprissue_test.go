package main

import "testing"

type areatest struct {
	arg1, expected float64
}

var areatests = []areatest{
	areatest{2,12.566371},
	areatest{3, 28.274334},
}
func TestArea(t *testing.T){
	for _,test := range areatests {
		if output := Area()
	}
}
