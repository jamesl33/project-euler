package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func loadNames() []string {
	var names []string

	file, err := os.Open("p022_names.txt")
	if err != nil {
		log.Fatalf("failed to open 'p022_names.txt': %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, strings.Split(scanner.Text(), ",")...)
	}

	return names
}

func main() {
	var (
		totalScore int
		names      = loadNames()
	)

	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	for index, name := range names {
		nameScore := 0

		for _, char := range name {
			if char != rune(34) {
				nameScore += int(char - 64)
			}
		}

		nameScore *= (index + 1)
		totalScore += nameScore
	}

	fmt.Println(totalScore)
}
