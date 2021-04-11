package main

import "fmt"

func cycleLength(n int) int {
	var (
		numbers = make([]int, n)
		divisor = 1
		index   int
	)

	for numbers[divisor] == 0 && divisor != 0 {
		numbers[divisor] = index
		divisor = (divisor * 10) % n
		index++
	}

	return index - numbers[divisor]
}

func main() {
	largest := 0

	for i := 1000; i > 1; i-- {
		length := cycleLength(i)

		if length > largest {
			largest = length
		}
	}

	fmt.Println(largest + 1)
}
