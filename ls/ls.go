package main

import "fmt"

type Phone interface {
	call()
}

type IPhone struct {
	Phone
	name  string
	model string
}

//
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

//
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

//// extend DroneX
type IPhone13Pro struct {
	IPhone
}

//
func (ph *IPhone13Pro) fly() {
	fmt.Printf(">>> call phone[%s] ... \n", ph.name)
	ph.boot()
	ph.launchKeypad()
	ph.dialing()
	ph.videoCall()
}

//
func (dr *IPhone13Pro) videoCall() {
	fmt.Println("[call] video call")
}

//
func Phones() []Phone {
	return []Phone{&IPhone{name: "Iphone13"}, &IPhone13Pro{IPhone{name: "Iphone13Pro"}}}
}

func main() {
	for _, ph := range Phones() {
		ph.call()
	}
}
