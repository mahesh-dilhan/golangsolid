package is

type Phone interface {
	call()
}

type IPhone struct {
	Phone
	name string
}
