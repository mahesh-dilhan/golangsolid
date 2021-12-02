package ocp

import "fmt"

type Phone interface {
	call()
}
type basicphone struct {
	name  string
	model string
}
type iPhone struct {
	Phone
	basicphone
}

func (p *iPhone) call() {
	fmt.Println("i'm ", p.name)
	fmt.Println("calling....")
	if p.model == "tesla" {
		fmt.Println("sattelite calling....")
	}
}

func main() {

}
