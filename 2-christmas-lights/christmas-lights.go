package main

const n = 1000
var lights [n][n]bool
var openLights = 0

func TurnOn(x1, y1, x2, y2 int) int{

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if !lights[i][j] {
				openLights++
				lights[i][j] = true
			}
		}
	}
	return openLights
}

func TurnOff(x1, y1, x2, y2 int) int{

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if lights[i][j] {
				openLights--
				lights[i][j] = false
			}
		}
	}
	return openLights
}