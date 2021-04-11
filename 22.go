package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func LoadNames(fileName string) []string {
	var names []string

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, strings.Split(scanner.Text(), ",")...)
	}

	return names
}

func main() {
	totalScore := 0
	names := LoadNames("names.txt")
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
