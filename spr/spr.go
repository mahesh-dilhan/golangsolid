package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type shape interface {
	name() string
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}
type square struct {
	length float64
}

func (r rectangle) name() string {
	return "rectangle"
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (s square) name() string {
	return "square"
}
func (s square) area() float64 {
	return s.length * s.length
}

type output struct{}

func (o output) Text(sh shape) string {
	return fmt.Sprintf("area of shape %s is %f", sh.name(), sh.area())
}

func NewSquare() square {
	s := &square{length: 2}
	return *s
}
func NewRectangle() rectangle {
	r := &rectangle{width: 2, height: 3}
	return *r
}

func (o output) JSON(sh shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	r := NewRectangle()
	s := NewSquare()

	o := output{}
	fmt.Println(o.Text(r))
	fmt.Println(o.JSON(r))

	fmt.Println(o.Text(s))
	fmt.Println(o.JSON(s))
}
