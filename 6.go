package main

import "fmt"

func main() {
	sum := 0
	sum_squares := 0

	for i := 1; i <= 100; i++ {
		sum += i
		sum_squares += i * i
	}

	fmt.Println((sum * sum) - sum_squares)
}
