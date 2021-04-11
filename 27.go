package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}

func main() {
	lcp := 0
	maxA := 0
	maxB := 0

	for a := -1000; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			cp := 0

			for isPrime((cp * cp) + (a * cp) + b) {
				cp++
			}

			if cp > lcp {
				lcp = cp
				maxA = a
				maxB = b
			}
		}
	}

	fmt.Println(maxA * maxB)
}
