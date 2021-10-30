package main

import "testing"

func TestChristmasLights(t *testing.T)  {

	x1 := 0
	y1 := 0
	x2 := 4
	y2 := 4
	
	got := TurnOn(x1, y1, x2, y2)
	want := 25

	if got != want{
		t.Errorf("want '%d' but got '%d' ", want, got)
	}
}