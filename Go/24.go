/* This file is part of project-euler

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

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
