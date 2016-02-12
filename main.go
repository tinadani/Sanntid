package main

import(
	"fmt"
	"./driver"
)

var buttonType = map[string]int{
	"Button call up": 0,
	"Button call down": 1,
    "Button internal panel": 2,
}

func main(){
	driver.InitializeElevator()
	fmt.Println(driver.ElevatorGetFloorSensorSignal())
	driver.ElevatorSetMotorDirection(0)
}
