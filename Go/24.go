package main

import (
    "fmt"
    "strings"
    "sort"
)

func GenPermutations(slice []string) [][]string {
    var helper func(slice []string, n int)
    var permutations [][]string

    helper = func(slice []string, n int) {
        if n == 1 {
            tmp := make([]string, len(slice))
            copy(tmp, slice)
            permutations = append(permutations, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(slice, n - 1)

                if n % 2 == 0 {
                    slice[i], slice[n - 1] = slice[n - 1], slice[i]
                } else {
                    slice[0], slice[n - 1] = slice[n - 1], slice[0]
                }
            }
        }
    }

    helper(slice, len(slice))

    return permutations
}

func main() {
    permutations := GenPermutations([]string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
    sort.Slice(permutations, func(i, j int) bool {return strings.Join(permutations[i], "") < strings.Join(permutations[j], "")})
    fmt.Println(strings.Join(permutations[1000000 - 1], ""))
}
