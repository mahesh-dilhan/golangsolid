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

func (ph *IPhone) call() {
	fmt.Printf(">>> call using phone[%s] ... \n", ph.name)
	ph.boot()
	ph.launchKeypad()
	ph.dialing()
}

func (ph *IPhone) launchKeypad() {
	fmt.Println("[launch] opening intent of keypad ")
}

func (ph *IPhone) dialing() {
	fmt.Println("[dial] connect to receiver  ")
}

func main() {
	var p1 Phone = &IPhone{name: "iPhone13"}
	p1.call()
	var p2 = &IPhone{name: "iPhone13 ProMax"}
	dr2.fly()
	dr2.add(&dr1)
}

func (p *IPhone) addSeries(l *Phone) {
	fmt.Printf("[lineup] setting up lineup, adding %v to series... \n", l)
}
