package ocp

type Phone interface {
	call()
}
type basicphone struct {
	name string
}
type iPhone struct {
	Phone
	basicphone
}

func (p *iPhone) call() {

}

func main() {

}
