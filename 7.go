package main

import (
	"fmt"
	"math"
)

func eratosthenesSieve(n int) []int {
	upperBound := 100

	if n > 6 {
		t := float64(n)
		upperBound = int(math.Ceil(t * (math.Log(t) + math.Log(math.Log(t)))))
	}

	sieve := make([]bool, upperBound)

	for i := 0; i < upperBound; i++ {
		sieve[i] = true
	}

	for i := 2; i < upperBound; i++ {
		if sieve[i] {
			for j := i * i; j < upperBound; j += i {
				sieve[j] = false
			}
		}
	}

	var primes []int

	for i := 0; i < upperBound; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes[2 : n+2]
}

func main() {
	fmt.Println(eratosthenesSieve(10001)[10000])
}
