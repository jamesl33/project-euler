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

func Contains(s string, c rune) bool {
	for _, e := range s {
		if e == c {
			return true
		}
	}

	return false
}

func IsPandigital(s string) bool {
	for i := 1; i <= 9; i++ {
		if !Contains(s, []rune(strconv.Itoa(i))[0]) {
			return false
		}
	}

	return true
}

func ConcatenatedProduct(n, c int) string {
	var products []int

	for i := 1; i <= c; i++ {
		products = append(products, n*i)
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(products)), ""), "[]")
}

func main() {
	largest := 0

	for i := 9000; i < 10000; i++ {
		c := ConcatenatedProduct(i, 2)

		if IsPandigital(c) {
			ci, _ := strconv.Atoi(c)

			if largest < ci {
				largest = ci
			}
		}
	}

	fmt.Println(largest)
}
