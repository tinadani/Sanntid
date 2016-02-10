package driver 

const MOTOR_SPEED = 2800
const N_FLOORS int = 4
const N_BUTTONS int = 3

button_type := map[string] int {
	"Button call up" : 2,
	"Button internal panel" : 1,
	"Button call down" : 0,
}

 var lamp_channels_matrix = [N_FLOORS][N_BUTTON] int{
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
 } 

 var button_channel_matrix = [N_FLOORS][N_BUTTONS] int{
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func Init_elevator() bool{
	init_success bool = IO_init)()
	if!init_success{
		return False
	}

	for floor:=0; floor < N_FLOORS; floor++{
		for int button = 0; button < N_BUTTONS; button++{
			Set_button_lamp(button,floor,0)
		}
	}
	Set_stop_lamp(0);
    Set_door_open_lamp(0);
    Set_floor_indicator(0);
}

func Elevator_set_motor_direction(motor_direction int){
    
    if motor_direction == 0{		
        IO_write_analog(MOTOR, 0)
    }
    else if motor_direction > 0 { 	//Set direction up
        IO_clear_bit(MOTORDIR)
        IO_write_analog(MOTOR, MOTOR_SPEED)
    } 
    else if motor_direction < 0{ 		//Set direction down
        IO_set_bit(MOTORDIR)
        IO_write_analog(MOTOR, MOTOR_SPEED)
    }
}

func Elevator_set_button_lamp(button_type int, floor int, on bool){

	if floor =< 0 || floor > N_FLOORS{
		fmt.Println('Invalid floor to set buttonlamp')
	}
	if 0 < button_type || button_type < N_BUTTONS {
		fmt.Println('Invalid button type')
	}
	if on {
		IO_set_bit(lamp_channels_matrix[floor][button_type])
	}
	else {
		IO_clear_bit(lamp_channels_matrix[floor][button_type])
	}
}

func Elevator_set_floor_indicator(floor int){
	 if floor =< 0 || floor > N_FLOORS{
		fmt.Println('Invalid floor to set floor indicator')
	}

    // Binary encoding. One light must always be on.
    if (floor & 0x02) {
        io_set_bit(LIGHT_FLOOR_IND1);
    } else {
        io_clear_bit(LIGHT_FLOOR_IND1);
    }    

    if (floor & 0x01) {
        io_set_bit(LIGHT_FLOOR_IND2);
    } else {
        io_clear_bit(LIGHT_FLOOR_IND2);
    }   
}

func Elevator_set_door_open_lamp(value int) {
    if (value) {
        io_set_bit(LIGHT_DOOR_OPEN)
    } 
    else {
        io_clear_bit(LIGHT_DOOR_OPEN)
    }
}

func Elevator_get_button_signal(button_type int, floor int) int{

	if floor =< 0 || floor > N_FLOORS{
		fmt.Println('Invalid floor to get button signal')
	}
	if 0 < button_type || button_type < N_BUTTONS {
		fmt.Println('Invalid button type to get button signal')
	}


    if io_read_bit(button_channels_matrix[floor][button_type]) {
        return 1
    }
    else {
        return 0
    }    
}

func Elevator_get_floor_sensor_signal() int{
    
    if IO_read_bit(SENSOR_FLOOR1) {
        return 0;
    } 
    else if IO_read_bit(SENSOR_FLOOR2) {
        return 1;
    } 
    else if IO_read_bit(SENSOR_FLOOR3) {
        return 2;
    } 
    else if IO_read_bit(SENSOR_FLOOR4) {
        return 3;
    } 
    else {
        return -1;
    }
}
