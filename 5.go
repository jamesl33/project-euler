package main

import "fmt"

func divisible(n int) bool {
	for i := 20; i > 0; i-- {
		if !(n%i == 0) {
			return false
		}
	}

	return true
}

func main() {
	current := 2520

	for true {
		if divisible(current) {
			fmt.Println(current)
			return
		}

		current++
	}
}
