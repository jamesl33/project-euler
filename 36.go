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

func Reverse(s string) string {
	c := []rune(s)

	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		c[left], c[right] = c[right], c[left]
	}

	return string(c)
}

func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

func main() {
	sum := 0

	for i := 0; i < 1000000; i++ {
		b2, b10 := fmt.Sprintf("%b", i), strconv.Itoa(i)

		if IsPalindrome(b2) && IsPalindrome(b10) {
			sum += i
		}
	}

	fmt.Println(sum)
}
