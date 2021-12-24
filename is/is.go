package is

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
