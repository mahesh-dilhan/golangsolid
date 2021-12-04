package main

import "fmt"

type tesla struct {
	Phone
	basicphone
}

func (p *tesla) hassattelitefeature() {
	fmt.Println("calling sattlete....")

}

func main() {
	p := tesla{}
	p.model = "iPhone"
	p.name = "iPhone 13"
	p.hassattelitefeature()
}
