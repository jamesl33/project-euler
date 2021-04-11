package main

import (
	"fmt"
	"math/big"
)

func main() {
	var powers []big.Int

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			var result big.Int
			result.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			powers = append(powers, result)
		}
	}

	unique := make(map[string]struct{})

	for _, p := range powers {
		if _, ok := unique[p.String()]; ok {
			continue
		}

		unique[p.String()] = struct{}{}
	}

	fmt.Println(len(unique))
}
