package main

const n = 1000
var lights [n][n]bool
var openLights = 0

func TurnOn(x1, y1, x2, y2 int) int{

	maxX, minX, maxY, minY := findMinMax(x1, x2, y1, y2)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if !lights[i][j]{
				openLights++
				lights[i][j] = true
			}
		}
	}
	return openLights
}

func TurnOff(x1, y1, x2, y2 int) int{

	maxX, minX, maxY, minY := findMinMax(x1, x2, y1, y2)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if lights[i][j] {
				openLights--
				lights[i][j] = false
			}
		}
	}
	return openLights
}

func Toggle(x1, y1, x2, y2 int) int{

	maxX, minX, maxY, minY := findMinMax(x1, x2, y1, y2)

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			
			if lights[i][j] {
				openLights--
				lights[i][j] = false
			}else {
				openLights++
				lights[i][j] = true
			}
		}
	}
	return openLights
}

func findMinMax(x1, x2, y1, y2 int) (maxX, minX, maxY, minY int){
	maxX = x1
	minX = x2
	maxY = y1
	minY = y2

	if maxX < minX{
		tempX := maxX
		maxX = minX
		minX = tempX
	}

	if maxY < minY{
		tempY := maxY
		maxY = minY
		minY = tempY
	}
	return
}