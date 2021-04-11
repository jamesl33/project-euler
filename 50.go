package main

import (
	"fmt"
)

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

func eratosthenesSieve(lowerBound, upperBound int) []int {
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

	for i := 2; i < upperBound; i++ {
		if sieve[i] && i > lowerBound {
			primes = append(primes, i)
		}
	}

	return primes
}

func contains(s []int, e int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return true
		}
	}

	return false
}

func main() {
	var (
		primes    = eratosthenesSieve(0, 1000000)
		sumPrimes = append(primes[:0:0], primes...)
	)

	for i := 1; i < len(sumPrimes); i++ {
		sumPrimes[i] = sumPrimes[i-1] + primes[i]
	}

	var largestConsecutive, largestPrime int

	for i := 0; i < len(primes); i++ {
		if i > len(primes)-largestConsecutive {
			break
		}

		for j := i; j < len(primes); j++ {
			sum := sumPrimes[j] - sumPrimes[i]

			if sum > 1000000 {
				break
			}

			if !isPrime(sum) {
				continue
			}

			if j-i < largestConsecutive {
				continue
			}

			largestConsecutive = j - i
			largestPrime = sum
		}
	}

	fmt.Println(largestPrime)
}
