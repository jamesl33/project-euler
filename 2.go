package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	total := 0
	fib_seq := 0
	fib_num := 0

	for fib_num < 4000000 {
		fib_num = fib(fib_seq)

		if fib_num%2 == 0 {
			total += fib_num
		}

		fib_seq++
	}

	fmt.Println(total)
}
