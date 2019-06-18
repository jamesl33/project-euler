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
    "sort"
    "strconv"
)

func EratosthenesSieve(lowerBound, upperBound int) []int {
    sieve := make([]bool, upperBound)

    for i := 2; i < int(math.Sqrt(float64(upperBound))); i++ {
        if !sieve[i] {
            for j := i * i; j < upperBound; j += i {
                sieve[j] = true
            }
        }
    }

    var primes []int

    for i := 2; i < upperBound; i++ {
        if !sieve[i] && i > lowerBound {
            primes = append(primes, i)
        }
    }

    return primes
}

func ArePermutations(s... string) bool {
    var sorted []string

    for _, str := range s {
        runeStr := []rune(str)

        sort.Slice(runeStr, func(i, j int) bool {
            return runeStr[i] < runeStr[j]
        })

        sorted = append(sorted, string(runeStr))
    }

    for _, str := range sorted {
        if str != sorted[0] {
            return false
        }
    }

    return true
}

func Contains(s []int, n int) bool {
    for _, e := range s {
        if e == n {
            return true
        }
    }

    return false
}

func main() {
    primes := EratosthenesSieve(1487, 10000)

    for i := 0; i < len(primes); i++ {
        for j := i + 1; j < len(primes); j++ {
            i, j, k := primes[i], primes[j], primes[j] + (primes[j] - primes[i])

            if k > 10000 || !Contains(primes, k) {
                continue
            }

            si, sj, sk := strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k)

            if !ArePermutations(si, sj, sk) {
                continue
            }

            fmt.Println(si + sj + sk)

            return
        }
    }
}
