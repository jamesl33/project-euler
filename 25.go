package main

import (
	"fmt"
	"math/big"
)

func FibSequence(n int) []*big.Int {
	sequence := []*big.Int{big.NewInt(0), big.NewInt(1)}

	for i := 2; i < n; i++ {
		var nextItem big.Int

		sequence = append(sequence, nextItem.Add(sequence[i-1], sequence[i-2]))
	}

	return sequence
}

func main() {
	current := 0

	for true {
		sequence := FibSequence(current)

		if len(sequence[len(sequence)-1].String()) == 1000 {
			fmt.Println(len(sequence) - 1)
			break
		}

		current++
	}
}
