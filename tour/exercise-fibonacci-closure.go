// Let's have some fun with functions.
// Implement a fibonacci function that returns a function
// (a closure) that returns successive fibonacci numbers
// (0, 1, 1, 2, 3, 5, ...).
package main

import "fmt"

// closures are functions that reference variables
// from outside their bodies
func fibonacci() func() int {
	firstNumber := -1
	secondNumber := -1
	return func() int {
		if firstNumber == -1 {
			firstNumber = secondNumber
			secondNumber++
			return secondNumber
		sum := firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
