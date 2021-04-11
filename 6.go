package main

import "fmt"

func main() {
	var sum, sumSquares int

	for i := 1; i <= 100; i++ {
		sum += i
		sumSquares += i * i
	}

	fmt.Println((sum * sum) - sumSquares)
}
