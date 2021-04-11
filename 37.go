package main

import (
	"fmt"
	"strconv"
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

func isLeftTruncatablePrime(s string) bool {
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i:])

		if !isPrime(n) {
			return false
		}
	}

	return true
}

func isRightTruncatablePrime(s string) bool {
	for i := len(s); i > 0; i-- {
		n, _ := strconv.Atoi(s[0:i])

		if !isPrime(n) {
			return false
		}
	}

	return true
}

func isTruncatablePrime(n int) bool {
	return isLeftTruncatablePrime(strconv.Itoa(n)) && isRightTruncatablePrime(strconv.Itoa(n))
}

func main() {
	var (
		sum, count int
		current    = 8
	)

	for count != 11 {
		if isTruncatablePrime(current) {
			sum += current
			count++
		}

		current++
	}

	fmt.Println(sum)
}
