package main

import "testing"

func TestChristmasLights(t *testing.T){

	//turn on
	//turn off
	//toggle
	//multiple instruction

	x1 := 0
	y1 := 0
	x2 := 4
	y2 := 4

	t.Run("turn on lights", func(t *testing.T) {

		got := TurnOn(x1, y1, x2, y2)
		want := 25

		if got != want{
			t.Errorf("want '%d' but got '%d' ", want, got)
		}
	})

	t.Run("turn off lights", func(t *testing.T) {

		got := TurnOff(x1, y1, x2, y2)
		want := 0

		if got != want{
			t.Errorf("want '%d' but got '%d' ", want, got)
		}
	})

	t.Run("multiple instruction", func(t *testing.T){

		got := TurnOn(0,0,99,99)
		got = TurnOn(89,99,50,10)
		got = TurnOff(29,0,20,49)
		got = Toggle(25,0,29,49)

		want := 9750

		if got != want{
			t.Errorf("want '%d' but got '%d' ", want, got)
		}
	})
}