package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPandigital(s string, l int) bool {
	for i := 1; i <= l; i++ {
		if !strings.ContainsRune(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

func main() {
	known := make(map[int]struct{})

	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			c := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i*j)

			if len(c) > 9 {
				break
			}

			if !isPandigital(c, 9) {
				continue
			}

			if _, ok := known[i*j]; ok {
				continue
			}

			known[i*j] = struct{}{}
		}
	}

	var sum int
	for k := range known {
		sum += k
	}

	fmt.Println(sum)
}
