package main

import (
	"fmt"
	"strconv"
)

func IsPrime(n int) bool {
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

func IsLeftTruncatablePrime(s string) bool {
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i:])

		if !IsPrime(n) {
			return false
		}
	}

	return true
}

func IsRightTruncatablePrime(s string) bool {
	for i := len(s); i > 0; i-- {
		n, _ := strconv.Atoi(s[0:i])

		if !IsPrime(n) {
			return false
		}
	}

	return true
}

func IsTruncatablePrime(n int) bool {
	return IsLeftTruncatablePrime(strconv.Itoa(n)) && IsRightTruncatablePrime(strconv.Itoa(n))
}

func main() {
	sum, count, current := 0, 0, 8

	for count != 11 {
		if IsTruncatablePrime(current) {
			sum += current
			count++
		}

		current++
	}

	fmt.Println(sum)
}
