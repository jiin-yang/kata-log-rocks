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

		got := TurnOn(0,0,100,100)
		got += TurnOn(90,100,50,10)
		got += TurnOff(30,0,20,50)
		got += Toggle(25,0,30,50)

		want := 9750

		if got != want{
			t.Errorf("want '%d' but got '%d' ", want, got)
		}
	})
}