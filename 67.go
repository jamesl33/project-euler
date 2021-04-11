package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readTriangle() [][]int {
	file, err := os.Open("p067_triangle.txt")
	if err != nil {
		log.Fatalf("failed to open 'p067_triangle.txt'")
	}
	defer file.Close()

	var (
		triangle = make([][]int, 0)
		scanner  = bufio.NewScanner(file)
	)

	for scanner.Scan() {
		row := make([]int, 0)

		for _, c := range strings.Split(scanner.Text(), " ") {
			n, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse '%s': %v", c, err)
			}

			row = append(row, int(n))
		}

		triangle = append(triangle, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to process number: %v", err)
	}

	return triangle
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	triangle := readTriangle()

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	fmt.Println(triangle[0][0])
}
