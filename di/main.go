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

func NewIphone(name string, s *Storage) *Phone {

}
