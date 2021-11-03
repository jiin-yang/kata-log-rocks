package main

var Area [9]string

func Move(number int, char string){
	Area[number-1] = char
}