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

func IsPrime(n int) bool {
    if n <= 1 {
        return false
    } else if n <= 3 {
        return true
    } else if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    i := 5

    for i * i <= n {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }

        i += 6
    }

    return true
}

func Reverse(slice *[]string, begin int, end int) {
    for left, right := begin, end; left < right; left, right = left+1, right-1 {
	(*slice)[left], (*slice)[right] = (*slice)[right], (*slice)[left]
    }
}

func Rotate(n int) func() int {
    slice := strings.Split(strconv.Itoa(n), "")

    return func() int {
        Reverse(&slice, 0, 0)
        Reverse(&slice, 1, len(slice) - 1)
        Reverse(&slice, 0, len(slice) - 1)

        rotation, _ := strconv.Atoi(strings.Join(slice, ""))
        return rotation
    }
}

func main() {
    count := 0

    for i := 0; i < 1000000; i++ {
        is_circular_prime := true
        rotate := Rotate(i)

        for r := 0; r < len(strconv.Itoa(i)); r++ {
            if !IsPrime(rotate()) {
                is_circular_prime = false
                break
            }
        }

        if is_circular_prime {
            count++
        }
    }

    fmt.Println(count)
}
