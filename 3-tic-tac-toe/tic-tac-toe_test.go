package main

import "testing"

func TestGameArea(t *testing.T){

	t.Run("first move", func(t *testing.T) {
		Move(2, "X")
		got := Area[1]
		want := "X"

		assertError(t, got, want)
	})

	t.Run("move to the same place", func(t *testing.T) {
		Move(2, "O")
		got := Area[1]
		want := "X"

		assertError(t, got, want)
	})

}

func TestWinCheck(t *testing.T){
	
	moves := [9]int{1,4,2,5,3,6,7,8,9}

	for i := 0; i < 9; i++ {
		t.Run("same line check", func(t *testing.T) {
			if i % 2 == 0{
				Move(moves[i], "X")
			}else {
				Move(moves[i], "O")
			}
			
			if i == 4{
				got := WinCheck()
				want := "X"

				assertError(t, got, want)
			}
		})
	}
}

func assertError(tb testing.TB, got, want string){
	tb.Helper()

	if got != want {
		tb.Errorf("want '%s' but got '%s' ", want, got)
	}
}
