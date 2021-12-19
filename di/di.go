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
	p := IPhone{name: "iPhone13", s: s, Phone: nil}
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
