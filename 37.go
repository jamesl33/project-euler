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

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}

func IsLeftTruncatablePrime(s string) bool {
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i:])

		if !IsPrime(n) {
			return false
		}
	}

	return true
}

func IsRightTruncatablePrime(s string) bool {
	for i := len(s); i > 0; i-- {
		n, _ := strconv.Atoi(s[0:i])

		if !IsPrime(n) {
			return false
		}
	}

	return true
}

func IsTruncatablePrime(n int) bool {
	return IsLeftTruncatablePrime(strconv.Itoa(n)) && IsRightTruncatablePrime(strconv.Itoa(n))
}

func main() {
	sum, count, current := 0, 0, 8

	for count != 11 {
		if IsTruncatablePrime(current) {
			sum += current
			count++
		}

		current++
	}

	fmt.Println(sum)
}
