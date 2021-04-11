package main

import "fmt"

func main() {
	combinations := []int{1}
	combinations = append(combinations, make([]int, 200)...)

	for _, coin := range []int{1, 2, 5, 10, 20, 50, 100, 200} {
		for index := coin; index < 201; index++ {
			combinations[index] += combinations[index-coin]
		}
	}

	fmt.Println(combinations[200])
}
