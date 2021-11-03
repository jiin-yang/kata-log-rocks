package main

import "testing"

func TestGameArea(t *testing.T){

	Move(2)
	got := Area[1]
	want := "X"

	if got != want {
		t.Errorf("want '%s' but got '%s' ", want, got)
	}

}
