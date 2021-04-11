package main

import (
	"fmt"
	"math"
)

type tuple struct {
	a, b int
}

func divisors(n int) []int {
	var divisors []int
	limit := int(math.Sqrt(float64(n)))

	for i := 1; i < limit+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i, n/i)
		}
	}

	for index, item := range divisors {
		if item == n {
			divisors = append(divisors[:index], divisors[index+1:]...)
			break
		}
	}

	return divisors
}

func sum(s []int) int {
	total := 0

	for _, item := range s {
		total += item
	}

	return total
}

func generateAmicablePairs(n int) []tuple {
	var amicablePairs []tuple

	for x := 1; x <= n; x++ {
		y := sum(divisors(x))

		if y > x && x == sum(divisors(y)) {
			amicablePairs = append(amicablePairs, tuple{x, y})
		}
	}

	return amicablePairs
}

func main() {
	total := 0

	for _, tuple := range generateAmicablePairs(10000) {
		total += tuple.a
		total += tuple.b
	}

	fmt.Println(total)
}
