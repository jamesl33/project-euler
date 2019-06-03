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

import "fmt"

func cycle_length(n int) int {
    numbers := make([]int, n)
    divisor := 1
    index := 0

    for numbers[divisor] == 0 && divisor != 0 {
        numbers[divisor] = index
        divisor = (divisor * 10) % n
        index++
    }

    return index - numbers[divisor]
}

func main() {
    largest_cycle := 0

    for i := 1000; i > 1; i-- {
        cycle_len := cycle_length(i)

        if cycle_len > largest_cycle {
            largest_cycle = cycle_len
        }
    }

    fmt.Println(largest_cycle + 1)
}