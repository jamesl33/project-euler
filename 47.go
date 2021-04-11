package main

import (
	"fmt"
)

func DistinctPrimeFactors(n int) []int {
	primeStatus := make([]bool, n+1)
	factorCount := make([]int, n+1)

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
	distinctPrimeFactors := DistinctPrimeFactors(1000000)

	count := 0

	for index, num := range distinctPrimeFactors {
		if num == 4 {
			count++

			if count == 4 {
				fmt.Println(index - 2)
				return
			}
		} else {
			count = 0
		}
	}
}
