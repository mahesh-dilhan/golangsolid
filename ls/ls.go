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
	fmt.Println("[boot..] do power on self test")
}

func (ph *IPhone) checkIO() {
	fmt.Println("[boot] check Input and output")
}

//
func (ph *IPhone) call() {
	fmt.Printf(">>> call using phone[%s] ... \n", ph.name)
	ph.boot()
	ph.launchKeypad()
	//ph.healthCheck()
}

func (ph *IPhone) launchKeypad() {
	fmt.Println("[launch] opening intent of keypad ")
}

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
