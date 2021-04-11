package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func readNumbers() []string {
	file, err := os.Open("p013_numbers.txt")
	if err != nil {
		log.Fatalf("failed to open 'p013_numbers.txt'")
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

	return lines
}

func main() {
	var (
		numbers = readNumbers()
		total   = new(big.Int)
	)

	for _, number := range numbers {
		num := new(big.Int)
		num.SetString(number, 10)
		total.Add(total, num)
	}

	fmt.Println(total.String()[:10])
}
