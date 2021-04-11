package main

import (
	"fmt"
	"sort"
	"strings"
)

func GenPermutations(str string) []string {
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

func main() {
	permutations := GenPermutations("0123456789")
	sort.Slice(permutations, func(i, j int) bool { return permutations[i] < permutations[j] })
	fmt.Println(permutations[1000000-1])
}
