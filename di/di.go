package main

import "fmt"

type Phone interface {
	call()
}

type Storage struct {
	sku string
}

type IPhone struct {
	Phone
	name string
	s    *Storage
}

func NewIphone(name string, s *Storage) *IPhone {
	p := IPhone{name: name, s: s, Phone: nil}
	return &p
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
	s1 := &Storage{sku: "256GB"}
	ip1 := NewIphone("iPhone13", s1)
	ip1.call()
	s2 := &Storage{sku: "1TB"}
	ip2 := NewIphone("iPhone13 Pro Max", s2)
	ip2.call()
}
