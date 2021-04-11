package main

import (
	"fmt"
	"math"
)

func isTriangle(n float64) bool {
	nthTriangle := (math.Sqrt(1+(8*n)) - 1) / 2
	return nthTriangle == math.Trunc(nthTriangle)
}

func isPentagonal(n float64) bool {
	nthPentagonal := (math.Sqrt(1+(24*n)) + 1) / 6
	return nthPentagonal == math.Trunc(nthPentagonal)
}

func isHexagonal(n float64) bool {
	nthHexagonal := (math.Sqrt(1+(8*n)) + 1) / 4
	return nthHexagonal == math.Trunc(nthHexagonal)
}

func main() {
	for i := 40756; i < 2147483647; i++ {
		if !isTriangle(float64(i)) || !isPentagonal(float64(i)) || !isHexagonal(float64(i)) {
			continue
		}

		fmt.Println(i)

		return
	}
}
