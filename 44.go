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

func IsPentagonal(n float64) bool {
    nthPentagonal := (math.Sqrt(1 + (24 * n)) + 1) / 6
    return nthPentagonal == math.Trunc(nthPentagonal)
}

func main() {
    for i := 1; ; i++ {
        n := i * ((3 * i) - 1) / 2

        for j := i - 1; j > 0; j-- {
            m := j * ((3 * j) - 1) / 2

            if !IsPentagonal(float64(n - m)) {
                continue
            }

            if !IsPentagonal(float64(n + m)) {
                continue
            }

            fmt.Println(n - m)

            return
        }
    }
}
