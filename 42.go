package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func readWords(filename string) []string {
	var words []string

	file, err := os.Open("p042_words.txt")
	if err != nil {
		log.Fatalf("failed to open 'p042_words.txt': %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, strings.Split(scanner.Text(), ",")...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to process words: %v", err)
	}

	return words
}

func isTriangleNum(n int) bool {
	nthTriangle := (math.Sqrt(float64(1+8*n)) - 1) / 2

	if nthTriangle == float64(int(nthTriangle)) {
		return true
	}

	return false
}

func isTriangleWord(s string) bool {
	sum := 0

	for _, c := range s {
		sum += int(c) - 96
	}

	return isTriangleNum(sum)
}

func main() {
	var count int

	for _, word := range readWords("p042_words.txt") {
		word = strings.ToLower(word)
		word = strings.Trim(word, "\"")

		if isTriangleWord(word) {
			count++
		}
	}

	fmt.Println(count)
}
