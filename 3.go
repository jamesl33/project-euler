package main

import "fmt"

func lpf(n int) int {
	var (
		largest int
		divisor = 2
	)

	for n > 1 {
		for n%divisor == 0 {
			largest = divisor
			n /= divisor
		}

		divisor++
	}

	return largest
}

func main() {
	fmt.Println(lpf(600851475143))
}
