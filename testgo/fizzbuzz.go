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

func fizzCheck(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return ""
}

func buzzCheck(n int) string {
	if n%5 == 0 {
		return "Buzz"
	}
	return ""
}

func FizzBuzz(n int) string {
	result := ""
	result += fizzCheck(n) + buzzCheck(n)
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}
