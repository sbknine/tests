package main

import "strconv"

func FizzBuzz(n int) string {
	mapFizz := map[bool]string{true: "Fizz"}
	mapBuzz := map[bool]string{true: "Buzz"}

	result := ""
	result += mapFizz[n%3 == 0] + mapBuzz[n%5 == 0]
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}
