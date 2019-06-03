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

func GenPermutations(str string) []string {
    var helper func(sli []string, ind int)

    var permutations []string

    helper = func(sli []string, ind int) {
        if ind == 1 {
            permutations = append(permutations, strings.Join(sli, ""))
        }

        for i := 0; i < ind; i++ {
            helper(sli, ind - 1)

            if ind % 2 == 0 {
                sli[i], sli[ind - 1] = sli[ind - 1], sli[i]
            } else {
                sli[0], sli[ind - 1] = sli[ind - 1], sli[0]
            }
        }
    }

    helper(strings.Split(str, ""), len(str))

    return permutations
}

func main() {
    permutations := GenPermutations("0123456789")
    sort.Slice(permutations, func(i, j int) bool {return permutations[i] < permutations[j]})
    fmt.Println(permutations[1000000 - 1])
}
