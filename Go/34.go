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
)

func Factorial(n int) int {
    f := []int {1}

    for i := 2; i <= n; i++ {
        f = append(f, f[(len(f) - 1)] * i)
    }

    return f[len(f) - 1]
}

func IsCurious(n int) bool {
    sum := 0

    for _, c := range strconv.Itoa(n) {
        ci, _ := strconv.Atoi(string(c))
        sum += Factorial(ci)
    }

    return sum == n
}

func main() {
    sum := 0

    for n := 3; n <= Factorial(9) * 7; n++ {
        if IsCurious(n) {
            sum += n
        }
    }

    fmt.Println(sum)
}
