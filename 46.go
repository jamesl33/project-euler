package main

import (
	"fmt"
	"math"
)

func EratosthenesSieve(n int) ([]int, []int) {
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

	var composites []int
	var primes []int

	for i := 0; i < upperBound; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}

		if !sieve[i] && i%2 != 0 {
			composites = append(composites, i)
		}
	}

	return composites, primes[2 : n+2]
}

func FitsGoldbach(n int, primes []int) bool {
	for _, prime := range primes {
		if prime >= n {
			return false
		}

		for num := 1; num < n/2; num++ {
			current := prime + (2 * (num * num))

			if current == n {
				return true
			}
		}
	}

	return false
}

func main() {
	composites, primes := EratosthenesSieve(10000)

	for _, composite := range composites {
		if !FitsGoldbach(composite, primes) {
			fmt.Println(composite)
			return
		}
	}
}
