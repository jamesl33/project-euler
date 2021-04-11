package main

import (
	"fmt"
	"strconv"
	"strings"
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
	for i := 1; i <= 9; i++ {
		if !contains(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

func concatenatedProduct(n, c int) string {
	var products []int

	for i := 1; i <= c; i++ {
		products = append(products, n*i)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(products)), ""), "[]")
}

func main() {
	largest := 0

	for i := 9000; i < 10000; i++ {
		c := concatenatedProduct(i, 2)

		if isPandigital(c) {
			ci, _ := strconv.Atoi(c)

			if largest < ci {
				largest = ci
			}
		}
	}

	fmt.Println(largest)
}
