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

type Tuple struct {
    a, b int
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

    return divisors
}

func Sum(s []int) int {
    total := 0

    for _, item := range s {
        total += item
    }

    return total
}

func GenAmicablePairs(n int) []Tuple {
    var amicablePairs []Tuple

    for x := 1; x <= n; x++ {
        y := Sum(Divisors(x))

        if y > x && x == Sum(Divisors(y)) {
            amicablePairs = append(amicablePairs, Tuple{x, y})
        }
    }

    return amicablePairs
}

func main() {
    total := 0

    for _, tuple := range GenAmicablePairs(10000) {
        total += tuple.a
        total += tuple.b
    }

    fmt.Println(total)
}
