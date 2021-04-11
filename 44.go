package main

import (
	"fmt"
	"math"
)

func isPentagonal(n float64) bool {
	nthPentagonal := (math.Sqrt(1+(24*n)) + 1) / 6
	return nthPentagonal == math.Trunc(nthPentagonal)
}

func main() {
	for i := 1; ; i++ {
		n := i * ((3 * i) - 1) / 2

		for j := i - 1; j > 0; j-- {
			m := j * ((3 * j) - 1) / 2

			if !isPentagonal(float64(n-m)) || !isPentagonal(float64(n+m)) {
				continue
			}

			fmt.Println(n - m)

			return
		}
	}
}
