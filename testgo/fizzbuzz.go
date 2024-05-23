package main

import "strconv"

/*
# FizzBuzz
กติกา
- commit ทุกครั้งที่ test ผ่าน
- each level do it two round at least
Level:
1. normal FizzBuzz.
2. FizzBuzz with ONE `if` no `switch` statement.
3. FizzBuzz with NO `if` and NO `switch` statement.
4. FizzBuzz with functional programming style.
5. FizzBuzz with Object-Oriented Programming (OOP) style.
*/

func FizzBuzz(n int) string {
	mapFizz := map[bool]string{true: "Fizz"}
	mapBuzz := map[bool]string{true: "Buzz"}
	mapNum := map[bool]string{false: strconv.Itoa(n)}

	result := ""
	result += mapFizz[n%3 == 0] + mapBuzz[n%5 == 0] + mapNum[n%3 == 0 || n%5 == 0]
	return result
}
