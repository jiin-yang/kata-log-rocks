package main

import (
"reflect"
"strconv"
"testing"
)

func TestFizzBuzz(t *testing.T){
	n := 15
	got := FizzBuzz(n)
	want := make([]string, n)

	for i := range want{
		j := i+1
		if j%15 == 0 {
			want[i] = "fizzBuzz"
		}else if j%3 == 0 {
			want[i]="fizz"
		}else if j%5 == 0 {
			want[i]="buzz"
		}else {
			want[i] = strconv.Itoa(j)
		}
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

