package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(int64(0))

	for i := 1; i <= 1000; i++ {
		next := big.NewInt(int64(i))
		next = next.Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
		sum = sum.Add(sum, next)
	}

	fmt.Println(sum.String()[len(sum.String())-10:])
}
