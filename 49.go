package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func EratosthenesSieve(lowerBound, upperBound int) []int {
	sieve := make([]bool, upperBound)

	for i := 2; i < int(math.Sqrt(float64(upperBound))); i++ {
		if !sieve[i] {
			for j := i * i; j < upperBound; j += i {
				sieve[j] = true
			}
		}
	}

	var primes []int

	for i := 2; i < upperBound; i++ {
		if !sieve[i] && i > lowerBound {
			primes = append(primes, i)
		}
	}

	return primes
}

func ArePermutations(s ...string) bool {
	var sorted []string

	for _, str := range s {
		runeStr := []rune(str)

		sort.Slice(runeStr, func(i, j int) bool {
			return runeStr[i] < runeStr[j]
		})

		sorted = append(sorted, string(runeStr))
	}

	for _, str := range sorted {
		if str != sorted[0] {
			return false
		}
	}

	return true
}

func Contains(s []int, n int) bool {
	for _, e := range s {
		if e == n {
			return true
		}
	}

	return false
}

func main() {
	primes := EratosthenesSieve(1487, 10000)

	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			i, j, k := primes[i], primes[j], primes[j]+(primes[j]-primes[i])

			if k > 10000 || !Contains(primes, k) {
				continue
			}

			si, sj, sk := strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k)

			if !ArePermutations(si, sj, sk) {
				continue
			}

			fmt.Println(si + sj + sk)

			return
		}
	}
}
