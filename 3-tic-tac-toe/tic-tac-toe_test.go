package main

import "testing"

func TestGameArea(t *testing.T){

	t.Run("first move", func(t *testing.T) {
		Move(2, "X")
		got := Area[1]
		want := "X"

		if got != want {
			t.Errorf("want '%s' but got '%s' ", want, got)
		}
	})

	t.Run("move to the same place", func(t *testing.T) {
		Move(2, "O")
		got := Area[1]
		want := "X"

		if got != want {
			t.Errorf("want '%s' but got '%s' ", want, got)
		}
	})

}
