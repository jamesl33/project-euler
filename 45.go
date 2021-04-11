package main

import (
	"fmt"
	"math"
)

func IsTriangle(n float64) bool {
	nthTriangle := (math.Sqrt(1+(8*n)) - 1) / 2
	return nthTriangle == math.Trunc(nthTriangle)
}

func IsPentagonal(n float64) bool {
	nthPentagonal := (math.Sqrt(1+(24*n)) + 1) / 6
	return nthPentagonal == math.Trunc(nthPentagonal)
}

func IsHexagonal(n float64) bool {
	nthHexagonal := (math.Sqrt(1+(8*n)) + 1) / 4
	return nthHexagonal == math.Trunc(nthHexagonal)
}

func main() {
	for i := 40756; i < 2147483647; i++ {
		if !IsTriangle(float64(i)) {
			continue
		}

		if !IsPentagonal(float64(i)) {
			continue
		}

		if !IsHexagonal(float64(i)) {
			continue
		}

		fmt.Println(i)

		return
	}
}
