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
    "strconv"
    "strings"
)

func IntContains(s []int, n int) bool {
    for _, e := range s {
        if e == n {
            return true
        }
    }

    return false
}

func IsPandigital(s string, l int) bool {
    for i := 1; i <= l; i++ {
        if !strings.ContainsRune(s, []rune(strconv.Itoa(i))[0]) {
            return false
        }
    }

    return true
}

func Sum(i ...int) int {
    sum := 0

    for e := range i {
        sum += i[e]
    }

    return sum
}

func main() {
    var k []int

    for i := 0; i < 10000; i++ {
        for j := 0; j < 10000; j++ {
            c := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(i * j)

            if len(c) == 9 && IsPandigital(c, 9) {
                if !IntContains(k, i * j) {
                    k = append(k, i * j)
                }
            } else if len(c) > 9 {
                break
            }
        }
    }

    fmt.Println(Sum(k...))
}
