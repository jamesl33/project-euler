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
)

func DistinctPrimeFactors(n int) []int {
    primeStatus := make([]bool, n + 1)
    factorCount := make([]int, n + 1)

    for i := 0; i < n + 1; i++ {
        primeStatus[i] = true
    }

    for i := 2; i < n + 1; i++ {
        if primeStatus[i] {
            factorCount[i]++

            for j := i * 2; j < n + 1; j += i {
                factorCount[j]++
                primeStatus[j] = false
            }
        }
    }

    return factorCount[1:]
}

func main() {
    distinctPrimeFactors := DistinctPrimeFactors(1000000)

    count := 0

    for index, num := range distinctPrimeFactors {
        if num == 4 {
            count++

            if count == 4 {
                fmt.Println(index - 2)
                return
            }
        } else {
            count = 0
        }
    }
}
