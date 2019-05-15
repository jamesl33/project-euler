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

func is_prime(n int) bool {
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

func main() {
    primes_found := 0
    last_prime := 0
    current_num := 0

    for primes_found <= 10000 {
        if is_prime(current_num) {
            primes_found++
            last_prime = current_num
        }

        current_num++
    }

    fmt.Println(last_prime)
}
