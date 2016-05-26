// Even Fibonacci numbers
// Problem 2
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import (
	"fmt"
)

const MAX int = 4000000

func main() {

	a, b, sum := 1, 1, 0

	for b < MAX {
		if b % 2 == 0 {
			sum += b
		}
		a, b = b, a + b
	}

	fmt.Println("Sum:", sum)
}