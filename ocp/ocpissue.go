package main

import "fmt"

type Phone interface {
	call()
}
type basicphone struct {
	name  string
	model string
}

func (iphone *basicphone) call() {
	fmt.Println("i'm ", iphone.name)
	fmt.Println("calling....")
	if iphone.model == "tesla" {
		fmt.Println("sattelite calling....")
	}
}

func main() {
	p := basicphone{}
	p.model = "iPhone"
	p.name = "iPhone 13"
	p.call()
}
