package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadWords(filename string) []string {
	var words []string

	if file, err := os.Open(filename); err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			words = append(words, strings.Split(scanner.Text(), ",")...)
		}
	}

	return words
}

func IsTriangleNum(n int) bool {
	nthTriangle := (math.Sqrt(float64(1+8*n)) - 1) / 2

	if nthTriangle == float64(int(nthTriangle)) {
		return true
	}

	return false
}

func IsTriangleWord(s string) bool {
	sum := 0

	for _, c := range s {
		sum += int(c) - 96
	}

	return IsTriangleNum(sum)
}

func main() {
	count := 0

	for _, word := range ReadWords("p042_words.txt") {
		word = strings.ToLower(word)
		word = strings.Trim(word, "\"")

		if IsTriangleWord(word) {
			count++
		}
	}

	fmt.Println(count)
}
