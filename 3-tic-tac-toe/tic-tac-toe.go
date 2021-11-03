package main

var Area [9]string

func Move(number int, char string){
	if Area[number-1] == ""{
		Area[number-1] = char
	}
}

func WinCheck() string{
	win := ""
	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			win = Area[0] + Area[1] + Area[2]
			break
		case 1:
			win = Area[3] + Area[4] + Area[5]
			break
		case 2:
			win = Area[6] + Area[7] + Area[8]
			break
		case 3:
			win = Area[0] + Area[4] + Area[8]
			break
		case 4:
			win = Area[0] + Area[3] + Area[6]
			break
		case 5:
			win = Area[1] + Area[4] + Area[7]
			break
		case 6:
			win = Area[2] + Area[5] + Area[8]
			break
		case 7:
			win = Area[2] + Area[4] + Area[6]
			break
		}
		if win == "XXX"{
			return "X"
		} else if(win == "OOO"){
			return "O"
		}
	}
	return ""
}