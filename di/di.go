package main

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
