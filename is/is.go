package is

import "fmt"

type Phone interface {
	call()
}

type IPhone struct {
	Phone
	name string
}

type Lineup interface {
	addSeries()
}

func (ph *IPhone) boot() {
	ph.doPOST()
	ph.checkIO()
}

func (ph *IPhone) doPOST() {
	fmt.Println("[boot] do power on self test")
}

func (ph *IPhone) checkIO() {
	fmt.Println("[boot] check Input and output")
}
