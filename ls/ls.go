package main

type Phone interface {
	call()
}

type IPhone struct {
	Phone
	name  string
	model string
}

//
//func (dr *DroneX) prepare() {
//	dr.checkBattery()
//	dr.checkPropeller()
//}
//
//func (dr *DroneX) fly() {
//	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
//	dr.prepare()
//	dr.takeOff()
//	dr.healthCheck()
//}
//
//func (dr *DroneX) checkBattery() {
//	fmt.Println("[preparing] checking battery's status ... ")
//}
//
//func (dr *DroneX) checkPropeller() {
//	fmt.Println("[preparing] checking propellers' status ... ")
//}
//
//func (dr *DroneX) takeOff() {
//	fmt.Println("[taking off] taking off now ... ")
//}
//
//func (dr *DroneX) healthCheck() {
//	fmt.Println("[flying] on the air, everything is ok, auto balancing enabled ... ")
//}
//
//// extend DroneX
//type DroneY struct {
//	DroneX
//}
//
//func (dr *DroneY) fly() {
//	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
//	dr.prepare()
//	dr.takeOff()
//	dr.healthCheck()
//	dr.spinAround()
//}
//
//func (dr *DroneY) spinAround() {
//	fmt.Println("[flying] I am spinning around ... ")
//}
//
//// extend DroneX
//type DroneZ struct {
//	DroneX
//}
//
//func (dr *DroneZ) fly() {
//	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
//	dr.prepare()
//	dr.takeOff()
//	dr.healthCheck()
//	dr.pirouettingCW()
//}
//
//func (dr *DroneZ) pirouettingCW() {
//	fmt.Println("[flying] I am pirouetting clockwise ... ")
//}
//
//func getDrones() []Drone {
//	return []Drone {&DroneY{}, &DroneZ{}, &DroneX{}}
//}
//
//func main() {
//	for _, dr := range getDrones() {
//		dr.fly()
//	}
//}
