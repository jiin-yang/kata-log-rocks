package main

var Area [9]string

func Move(number int, char string){
	if Area[number-1] == ""{
		Area[number-1] = char
	}
}