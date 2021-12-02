package ocp

type Phone interface {
	call()
}
type iPhone struct {
	Phone
	name string
}
type tsla struct {
	Phone
	name string
}

func main() {

}
