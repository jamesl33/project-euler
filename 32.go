package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IntContains(s []int, n int) bool {
	for _, e := range s {
		if e == n {
			return true
		}
	}

	return false
}

func IsPandigital(s string, l int) bool {
	for i := 1; i <= l; i++ {
		if !strings.ContainsRune(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

func Sum(i ...int) int {
	sum := 0

	for e := range i {
		sum += i[e]
	}

	return sum
}

func main() {
	var k []int

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			c := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i*j)

			if len(c) == 9 && IsPandigital(c, 9) {
				if !IntContains(k, i*j) {
					k = append(k, i*j)
				}
			} else if len(c) > 9 {
				break
			}
		}
	}

	fmt.Println(Sum(k...))
}
