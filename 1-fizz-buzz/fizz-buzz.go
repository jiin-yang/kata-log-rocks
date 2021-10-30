package main

import "strconv"

func FizzBuzz(n int) []string {

	result := make([]string, n)

	for i := 1; i <= n; i++{

		if i%15==0 {
			result[i-1] = "fizzBuzz"
		}else if i%3== 0 {
			result[i-1] = "fizz"
		}else if i%5== 0 {
			result[i-1] = "buzz"
		}else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

