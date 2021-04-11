package main

import (
	"fmt"
	"strconv"
)

func getDigit(i int) int {
	var (
		currentRange = 9
		first        = 1
		numDigits    = 1
		skip         int
	)

	for skip+currentRange*numDigits < i {
		skip += currentRange * numDigits
		currentRange *= 10
		first *= 10
		numDigits++
	}

	for currentRange > 9 {
		for skip+currentRange*numDigits < i {
			skip += currentRange * numDigits
			first += currentRange
		}

		currentRange /= 10
	}

	for skip+numDigits < i {
		skip += numDigits
		first++
	}

	num, _ := strconv.Atoi(string(strconv.Itoa(first)[i-(skip+1)]))

	return num
}

func main() {
	total := 1

	for i := 1; i <= 1000000; i *= 10 {
		total *= getDigit(i)
	}

	fmt.Println(total)
}
