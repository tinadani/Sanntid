package driver 

import (
    "fmt"
)

const MOTORSPEED = 2800
const NUMFLOORS int = 4
const NUMBUTTONS int = 3
const ON bool = true
const OFF bool = false
var buttonType = map[string]int{
	"Button call up": 0,
	"Button call down": 1,
    "Button internal panel": 2,
}

 var lampChannelsMatrix = [NUMFLOORS][NUMBUTTONS] int{
    {LIGHTUP1, LIGHTDOWN1, LIGHTCOMMAND1},
    {LIGHTUP2, LIGHTDOWN2, LIGHTCOMMAND2},
    {LIGHTUP3, LIGHTDOWN3, LIGHTCOMMAND3},
    {LIGHTUP4, LIGHTDOWN4, LIGHTCOMMAND4},
 } 

 var buttonChannelsMatrix = [NUMFLOORS][NUMBUTTONS] int{
    {BUTTONUP1, BUTTONDOWN1, BUTTONCOMMAND1},
    {BUTTONUP2, BUTTONDOWN2, BUTTONCOMMAND2},
    {BUTTONUP3, BUTTONDOWN3, BUTTONCOMMAND3},
    {BUTTONUP4, BUTTONDOWN4, BUTTONCOMMAND4},
}


func ElevatorSetMotorDirection(motorDirection int){

    if motorDirection > 0 { 	//Set direction up if positive number
        IOClearBit(MOTORDIR)
        IOWriteAnalog(MOTOR, MOTORSPEED)
    }else if motorDirection < 0 { 		//Set direction down if negative number
        IOSetBit(MOTORDIR)
        IOWriteAnalog(MOTOR, MOTORSPEED)
    }else if (motorDirection == 0){        
        IOWriteAnalog(MOTOR, 0)     //if not stop elevator
    }else{
        fmt.Println("Unable to set motor direction")
    }
}

func ElevatorSetButtonLamp(buttonType int, floor int, on bool){

	if (floor < 0) || (floor > NUMFLOORS){
		fmt.Println("Invalid floor to set buttonlamp")
	}

    if (0 > buttonType) || (buttonType > NUMBUTTONS) {
		fmt.Println("Invalid button type")
	}

    if on {
		IOSetBit(lampChannelsMatrix[floor][buttonType])
	}else {
		IOClearBit(lampChannelsMatrix[floor][buttonType])
	}
}

func ElevatorSetFloorIndicator(floor int){
	 if floor < 0 || floor > NUMFLOORS{
		fmt.Println("Invalid floor to set floor indicator")
	}

    // Binary encoding. One light must always be on.
    if (floor & 0x02 != 0) {
        IOSetBit(LIGHTFLOORIND1);
    } else {
        IOClearBit(LIGHTFLOORIND1);
    }    

    if (floor & 0x01 != 0) {
        IOSetBit(LIGHTFLOORIND2);
    } else {
        IOClearBit(LIGHTFLOORIND2);
    }   
}

func ElevatorSetDoorOpenLamp(on bool) {
    if (on) {
        IOSetBit(LIGHTDOOROPEN)
    }else {
        IOClearBit(LIGHTDOOROPEN)
    }
}

func ElevatorGetButtonSignal(buttonType int, floor int) int{

	if (floor < 0 )|| (floor > NUMFLOORS){
		fmt.Println("Invalid floor to get button signal")
	}
	if (0 >= buttonType) || (buttonType > NUMBUTTONS ){
		fmt.Println("Invalid button type to get button signal")
	}

    if IOReadBit(buttonChannelsMatrix[floor][buttonType]) !=0 {
        return 1
    }else {
        return 0
    }    
}

func ElevatorGetFloorSensorSignal() int{
    
    if IOReadBit(SENSORFLOOR1) != 0 {
        return 0;
    }else if IOReadBit(SENSORFLOOR2) != 0 {
        return 1;
    }else if IOReadBit(SENSORFLOOR3) != 0 { 
        return 2;
    }else if IOReadBit(SENSORFLOOR4) != 0 {
        return 3;
    }else {
        return -1;
    }
}


func InitializeElevator() bool{
    var initSuccess bool = IOInitializeElevator()
    
    if !initSuccess {
        return false
    }

    for floor := 0; floor < NUMFLOORS; floor++{
        for button := 0; button < NUMBUTTONS; button++{
            ElevatorSetButtonLamp(button,floor,OFF)
        }
    }

    ElevatorSetDoorOpenLamp(OFF)
    //ElevatorSetFloorIndicator(ElevatorGetFloorSensorSignal())
    return true
}
