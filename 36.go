package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	c := []rune(s)

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		c[left], c[right] = c[right], c[left]
	}

	return string(c)
}

func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

func main() {
	sum := 0

	for i := 0; i < 1000000; i++ {
		b2, b10 := fmt.Sprintf("%b", i), strconv.Itoa(i)

		if IsPalindrome(b2) && IsPalindrome(b10) {
			sum += i
		}
	}

	fmt.Println(sum)
}
