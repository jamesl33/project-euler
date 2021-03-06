package main

import "fmt"

func main() {
	collatz := [][]int{}

	for i := 1; i <= 1000000; i++ {
		var (
			n        = i
			sequence = make([]int, 0)
		)

		for n != 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = n*3 + 1
			}

			sequence = append(sequence, n)
		}

		collatz = append(collatz, sequence)
	}

	longest := 0

	for index, sequence := range collatz {
		if len(sequence) > len(collatz[longest]) {
			longest = index
		}
	}

	fmt.Println(longest + 1)
}
