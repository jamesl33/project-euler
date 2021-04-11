// My solution
// package main
//
// import "fmt"
//
// func product(slice []int) int {
// 	product := 1
//
// 	for _, num := range slice {
// 		product *= num
// 	}
//
// 	return product
// }
//
// func main() {
// 	numbers := [][]int{
// 		{8, 2, 22, 97, 38, 15, 00, 40, 00, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
// 		{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 00},
// 		{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65},
// 		{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91},
// 		{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
// 		{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
// 		{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
// 		{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
// 		{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
// 		{21, 36, 23, 9, 75, 00, 76, 44, 20, 45, 35, 14, 00, 61, 33, 97, 34, 31, 33, 95},
// 		{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92},
// 		{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 00, 17, 54, 24, 36, 29, 85, 57},
// 		{86, 56, 00, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
// 		{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40},
// 		{4, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
// 		{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
// 		{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
// 		{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16},
// 		{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54},
// 		{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
// 	}
//
// 	adjacent := 4
// 	largest_sum := 0
//
// 	for i, arr := range numbers {
// 		for j, _ := range arr {
// 			if i >= adjacent {
// 				var numbers_up []int
//
// 				for n := i - 1; n > (i-adjacent)-1; n-- {
// 					numbers_up = append(numbers_up, numbers[n][j])
// 				}
//
// 				sum_up := product(numbers_up)
//
// 				if sum_up > largest_sum {
// 					largest_sum = sum_up
// 				}
// 			}
//
// 			if i <= len(numbers)-adjacent {
// 				var numbers_down []int
//
// 				for n := i; n < i+adjacent; n++ {
// 					numbers_down = append(numbers_down, numbers[n][j])
// 				}
//
// 				sum_down := product(numbers_down)
//
// 				if sum_down > largest_sum {
// 					largest_sum = sum_down
// 				}
// 			}
//
// 			if j >= adjacent {
// 				sum_left := product(numbers[i][j-adjacent : j])
//
// 				if sum_left > largest_sum {
// 					largest_sum = sum_left
// 				}
// 			}
//
// 			if j <= len(arr)-adjacent {
// 				sum_right := product(numbers[i][j : j+adjacent])
//
// 				if sum_right > largest_sum {
// 					largest_sum = sum_right
// 				}
// 			}
//
// 			if i >= adjacent-1 && j >= adjacent-1 {
// 				var numbers_upper_left []int
//
// 				for n := 0; n <= adjacent-1; n++ {
// 					numbers_upper_left = append(numbers_upper_left, numbers[i-n][j-n])
// 				}
//
// 				sum_upper_left := product(numbers_upper_left)
//
// 				if sum_upper_left > largest_sum {
// 					largest_sum = sum_upper_left
// 				}
// 			}
//
// 			if i >= adjacent-1 && j <= len(arr)-adjacent {
// 				var numbers_upper_right []int
//
// 				for n := 0; n < adjacent; n++ {
// 					numbers_upper_right = append(numbers_upper_right, numbers[i-n][j+n])
// 				}
//
// 				sum_upper_right := product(numbers_upper_right)
//
// 				if sum_upper_right > largest_sum {
// 					largest_sum = sum_upper_right
// 				}
// 			}
//
// 			if i <= len(numbers)-adjacent && j <= len(arr)-adjacent {
// 				var numbers_lower_right []int
//
// 				for n := 0; n < adjacent; n++ {
// 					numbers_lower_right = append(numbers_lower_right, numbers[i+n][j+n])
// 				}
//
// 				sum_lower_right := product(numbers_lower_right)
//
// 				if sum_lower_right > largest_sum {
// 					largest_sum = sum_lower_right
// 				}
// 			}
//
// 			if i <= len(numbers)-adjacent && j >= adjacent-1 {
// 				var numbers_lower_left []int
//
// 				for n := 0; n < adjacent; n++ {
// 					numbers_lower_left = append(numbers_lower_left, numbers[i+n][j-n])
// 				}
//
// 				sum_lower_left := product(numbers_lower_left)
//
// 				if sum_lower_left > largest_sum {
// 					largest_sum = sum_lower_left
// 				}
// 			}
// 		}
// 	}
//
// 	fmt.Println(largest_sum)
// }

// Better solution adapted from solution found here - https://github.com/nayuki/Project-Euler-solutions/blob/master/python/p011.py
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readGrid() [][]int {
	file, err := os.Open("p011_grid.txt")
	if err != nil {
		log.Fatalf("failed to open 'p011_grid.txt'")
	}
	defer file.Close()

	var (
		grid    = make([][]int, 0)
		scanner = bufio.NewScanner(file)
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

		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to process number: %v", err)
	}

	return grid
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func gridProduct(array [][]int, startX int, startY int, dirX int, dirY int, adj int) int {
	result := 1

	for i := 0; i < adj; i++ {
		result *= array[startY+i*dirY][startX+i*dirX]
	}

	return result
}

func main() {
	var (
		grid   = readGrid()
		result = -1
		adj    = 4
		height = len(grid)
		width  = len(grid[0])
	)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x+adj <= width {
				result = max(gridProduct(grid, x, y, 1, 0, adj), result)
			}

			if x-adj >= width {
				result = max(gridProduct(grid, x, y, -1, 0, adj), result)
			}

			if y+adj <= height {
				result = max(gridProduct(grid, x, y, 0, 1, adj), result)
			}

			if y-adj >= 0 {
				result = max(gridProduct(grid, x, y, 0, -1, adj), result)
			}

			if x+adj <= width && y+adj <= height {
				result = max(gridProduct(grid, x, y, 1, 1, adj), result)
			}

			if x-adj >= 0 && y-adj >= 0 {
				result = max(gridProduct(grid, x, y, -1, -1, adj), result)
			}

			if x+adj <= height && y-adj >= 0 {
				result = max(gridProduct(grid, x, y, 1, -1, adj), result)
			}

			if x-adj >= 0 && y+adj <= height {
				result = max(gridProduct(grid, x, y, -1, 1, adj), result)
			}
		}
	}

	fmt.Println(result)
}
