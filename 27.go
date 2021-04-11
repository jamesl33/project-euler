package main

import "fmt"

func is_prime(n int) bool {
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
	longest_consecutive_primes := 0
	max_a := 0
	max_b := 0

	for a := -1000; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			consecutive_primes := 0

			for is_prime((consecutive_primes * consecutive_primes) + (a * consecutive_primes) + b) {
				consecutive_primes++
			}

			if consecutive_primes > longest_consecutive_primes {
				longest_consecutive_primes = consecutive_primes
				max_a = a
				max_b = b
			}
		}
	}

	fmt.Println(max_a * max_b)
}
