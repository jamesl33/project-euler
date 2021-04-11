package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	var total, fibSeq, fibNum int

	for fibNum < 4000000 {
		fibNum = fib(fibSeq)

		if fibNum%2 == 0 {
			total += fibNum
		}

		fibSeq++
	}

	fmt.Println(total)
}
