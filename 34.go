package main

import (
	"fmt"
	"strconv"
)

func factorial(n int) int {
	f := []int{1}

	for i := 2; i <= n; i++ {
		f = append(f, f[(len(f)-1)]*i)
	}

	return f[len(f)-1]
}

func isCurious(n int) bool {
	sum := 0

	for _, c := range strconv.Itoa(n) {
		ci, _ := strconv.Atoi(string(c))
		sum += factorial(ci)
	}

	return sum == n
}

func main() {
	sum := 0

	for n := 3; n <= factorial(9)*7; n++ {
		if isCurious(n) {
			sum += n
		}
	}

	fmt.Println(sum)
}
