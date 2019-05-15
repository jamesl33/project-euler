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
    "math"
)

func RemoveDuplicates(slice []int) []int {
    known := map[int]bool {}
    result := []int {}

    for _, value := range slice {
        if !known[value] {
            known[value] = true
            result = append(result, value)
        }
    }

    return result
}

func Divisors(n int) []int {
    var divisors []int
    limit := int(math.Sqrt(float64(n)))

    for i := 1; i < limit + 1; i++ {
        if n % i == 0 {
            divisors = append(divisors, i, n / i)
        }
    }

    for index, item := range divisors {
        if item == n {
            divisors = append(divisors[:index], divisors[index + 1:]...)
            break
        }
    }

    return RemoveDuplicates(divisors)
}

func Sum(s []int) int {
    total := 0

    for _, item := range s {
        total += item
    }

    return total
}

func GenAbundantNumbers() []int {
    var abundantNumbers []int

    for i := 0; i < 28123; i++ {
        if Sum(Divisors(i)) > i {
            abundantNumbers = append(abundantNumbers, i)
        }
    }

    return abundantNumbers
}

func main() {
    abundantNumbers := GenAbundantNumbers()
    var expressible [28123]bool

    for _, i := range abundantNumbers {
        for _, j := range abundantNumbers {
            if i + j < 28123 {
                expressible[i + j] = true
            } else {
                break
            }
        }
    }

    total := 0

    for index, value := range expressible {
        if !value {
            total += index
        }
    }

    fmt.Println(total)
}
