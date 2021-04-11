package main

import "fmt"

func cycle_length(n int) int {
	numbers := make([]int, n)
	divisor := 1
	index := 0

	for numbers[divisor] == 0 && divisor != 0 {
		numbers[divisor] = index
		divisor = (divisor * 10) % n
		index++
	}

	return index - numbers[divisor]
}

func main() {
	largest_cycle := 0

	for i := 1000; i > 1; i-- {
		cycle_len := cycle_length(i)

		if cycle_len > largest_cycle {
			largest_cycle = cycle_len
		}
	}

	fmt.Println(largest_cycle + 1)
}
