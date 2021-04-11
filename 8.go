package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readNumber() string {
	file, err := os.Open("p008_number.txt")
	if err != nil {
		log.Fatalf("failed to open 'p008_number.txt'")
	}
	defer file.Close()

	var (
		lines   = make([]string, 0)
		scanner = bufio.NewScanner(file)
	)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to process number: %v", err)
	}

	return strings.Join(lines, "")
}

func productSubstr(substr string) int {
	product := 1

	for _, char := range substr {
		product *= int(char - '0')
	}

	return product
}

func main() {
	var (
		largest int
		number  = readNumber()
	)

	for i := 0; i < len(number)-13+1; i++ {
		sum := productSubstr(number[i : i+13])

		if sum > largest {
			largest = sum
		}
	}

	fmt.Println(largest)
}
