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

func is_palindrome(n int) bool {
	original := n
	reversed := 0
	remainder := 0

	for n != 0 {
		remainder = n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}

	return original == reversed
}

func main() {
	largest := 0

	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			product := i * j

			if is_palindrome(product) && product > largest {
				largest = product
			}
		}
	}

	fmt.Println(largest)
}
