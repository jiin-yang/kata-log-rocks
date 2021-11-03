package main

import "testing"

func TestGameArea(t *testing.T){

	Move(2)
	got := Area[1]
	want := "X"

	if got != want {
		t.Errorf("want '%q' but got '%q' ", want, got)
	}

}
