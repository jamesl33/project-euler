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

func GetDigit(i int) int {
	currentRange := 9
	first := 1
	numDigits := 1
	skip := 0

	for skip+currentRange*numDigits < i {
		skip += currentRange * numDigits
		currentRange *= 10
		first *= 10
		numDigits++
	}

	for currentRange > 9 {
		for skip+currentRange*numDigits < i {
			skip += currentRange * numDigits
			first += currentRange
		}

		currentRange /= 10
	}

	for skip+numDigits < i {
		skip += numDigits
		first++
	}

	num, _ := strconv.Atoi(string(strconv.Itoa(first)[i-(skip+1)]))

	return num
}

func main() {
	total := 1

	for i := 1; i <= 1000000; i *= 10 {
		total *= GetDigit(i)
	}

	fmt.Println(total)
}
