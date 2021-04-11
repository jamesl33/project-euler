package main

import (
	"fmt"
	"math"
	"strconv"
)

func SumFifthPower(n int) int {
	sum := 0

	for _, c := range strconv.Itoa(n) {
		sum += int(math.Pow(float64(c-'0'), float64(5)))
	}

	return sum
}

func main() {
	sum := 0

	for i := 64; i < 295245; i++ {
		if i == SumFifthPower(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
