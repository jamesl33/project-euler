package main

import "fmt"

func isPalindrome(n int) bool {
	var (
		original            = n
		reversed, remainder int
	)

	for n != 0 {
		remainder = n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	return original == reversed
}

func main() {
	largest := 0

	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			product := i * j

			if isPalindrome(product) && product > largest {
				largest = product
			}
		}
	}

	fmt.Println(largest)
}
