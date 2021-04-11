package main

import (
	"fmt"
	"math/big"
)

func contains(slice []big.Int, n *big.Int) bool {
	for _, e := range slice {
		if e.Cmp(n) == 0 {
			return true
		}
	}

	return false
}

func main() {
	var powers []big.Int

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			var result big.Int
			result.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			powers = append(powers, result)
		}
	}

	var unique_powers []big.Int

	for _, p := range powers {
		if !contains(unique_powers, &p) {
			unique_powers = append(unique_powers, p)
		}
	}

	fmt.Println(len(unique_powers))
}
