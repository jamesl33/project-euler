package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func generatePermutations(str string) []string {
	var helper func(sli []string, ind int)

	var permutations []string

	helper = func(sli []string, ind int) {
		if ind == 1 {
			permutations = append(permutations, strings.Join(sli, ""))
		}

		for i := 0; i < ind; i++ {
			helper(sli, ind-1)

			if ind%2 == 0 {
				sli[i], sli[ind-1] = sli[ind-1], sli[i]
			} else {
				sli[0], sli[ind-1] = sli[ind-1], sli[0]
			}
		}
	}

	helper(strings.Split(str, ""), len(str))

	return permutations
}

func isPandigital(s string) bool {
	for i := 0; i <= 9; i++ {
		if !strings.ContainsRune(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

func isSpecialPandigital(s string) bool {
	primes := []int{2, 3, 5, 7, 11, 13, 17}

	for i := 1; i < 8; i++ {
		n, _ := strconv.Atoi(s[i : i+3])

		if !(n%primes[i-1] == 0) {
			return false
		}
	}

	return true
}

func main() {
	permutations := generatePermutations("0123456789")

	sum := &big.Int{}

	for _, permutation := range permutations {
		if permutation[0] != '0' && isPandigital(permutation) && isSpecialPandigital(permutation) {
			n := &big.Int{}
			n.SetString(permutation, 10)

			sum.Add(sum, n)
		}
	}

	fmt.Println(sum)
}
