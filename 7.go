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

func EratosthenesSieve(n int) []int {
    upperBound := 100

    if n > 6 {
        t := float64(n)
        upperBound = int(math.Ceil(t * (math.Log(t) + math.Log(math.Log(t)))))
    }

    sieve := make([]bool, upperBound)

    for i := 0; i < upperBound; i++ {
        sieve[i] = true
    }

    for i := 2; i < upperBound; i++ {
        if sieve[i] {
            for j := i * i; j < upperBound; j += i {
                sieve[j] = false
            }
        }
    }

    var primes []int

    for i := 0; i < upperBound; i++ {
        if sieve[i] {
            primes = append(primes, i)
        }
    }

    return primes[2:n + 2]
}

func main() {
    fmt.Println(EratosthenesSieve(10001)[10000])
}
