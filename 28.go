package main

import "fmt"

// func main() {
// 	// 3 13 31 = 4(n * n) - 2n + 1
// 	// 5 17 37 = 4(n * n) + 1
// 	// 7 21 43 = 4(n * n) + 2n + 1
// 	// 9 25 49 = 4(n * n) + 4n + 1
//
// 	sum := 0
//
// 	for i := 1; i <= (1001-1)/2; i++ {
// 		sum += 4*(i*i) - (2 * i) + 1
// 		sum += 4*(i*i) + 1
// 		sum += 4*(i*i) + (2 * i) + 1
// 		sum += 4*(i*i) + (4 * i) + 1
// 	}
//
// 	// add 1 because the sequences do not take into account the 1 in the centre
// 	fmt.Println(sum + 1)
// }

func main() {
	sum := 0

	for i := 1; i*i <= 1001*1001; i += 2 {
		for w := 0; w < 4; w++ {
			sum += (i * i) - ((i - 1) * w)
		}
	}

	// minus 3 because centre '1' will be processed 4 times overall
	fmt.Println(sum - 3)
}
