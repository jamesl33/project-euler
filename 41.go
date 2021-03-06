package main

import (
	"fmt"
	"strconv"
)

func contains(s string, c rune) bool {
	for _, e := range s {
		if e == c {
			return true
		}
	}

	return false
}

func isPandigital(s string) bool {
	for i := 1; i <= len(s); i++ {
		if !contains(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

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
	for i := 7654321; i > 0; i-- {
		if isPandigital(strconv.Itoa(i)) && isPrime(i) {
			fmt.Println(i)
			return
		}
	}
}
