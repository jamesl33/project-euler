package main

import "fmt"

func main() {
	max, maxp := 0, 0

	for p := 2; p <= 1000; p += 2 {
		count := 0

		for a := 2; a < p/3; a++ {
			if p*(p-2*a)%(2*(p-a)) == 0 {
				count++
			}
		}

		if count > max {
			max = count
			maxp = p
		}
	}

	fmt.Println(maxp)
}
