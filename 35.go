package main

import (
	"fmt"
	"strconv"
	"strings"
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

func reverse(slice []string, begin int, end int) {
	for left, right := begin, end; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
}

func rotate(n int) func() int {
	slice := strings.Split(strconv.Itoa(n), "")

	return func() int {
		reverse(slice, 0, 0)
		reverse(slice, 1, len(slice)-1)
		reverse(slice, 0, len(slice)-1)

		rotation, _ := strconv.Atoi(strings.Join(slice, ""))
		return rotation
	}
}

func main() {
	count := 0

	for i := 0; i < 1000000; i++ {
		is_circular_prime := true
		rotate := rotate(i)

		for r := 0; r < len(strconv.Itoa(i)); r++ {
			if !isPrime(rotate()) {
				is_circular_prime = false
				break
			}
		}

		if is_circular_prime {
			count++
		}
	}

	fmt.Println(count)
}
