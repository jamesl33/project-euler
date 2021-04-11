package main

import (
	"fmt"
)

func distinctPrimeFactors(n int) []int {
	var (
		primeStatus = make([]bool, n+1)
		factorCount = make([]int, n+1)
	)

	for i := 0; i < n+1; i++ {
		primeStatus[i] = true
	}

	for i := 2; i < n+1; i++ {
		if primeStatus[i] {
			factorCount[i]++

			for j := i * 2; j < n+1; j += i {
				factorCount[j]++
				primeStatus[j] = false
			}
		}
	}

	return factorCount[1:]
}

func main() {
	var (
		count        int
		primeFactors = distinctPrimeFactors(1000000)
	)

	for index, num := range primeFactors {
		if num != 4 {
			count = 0
			continue
		}

		count++

		if count == 4 {
			fmt.Println(index - 2)
			return
		}
	}
}
