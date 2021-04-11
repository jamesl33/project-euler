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

func Num2Words(n int) string {
	ones := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if 0 <= n && n < 20 {
		return ones[n]
	} else if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		} else {
			return tens[n/10] + ones[n%10]
		}
	} else if n < 1000 {
		if n%100 == 0 {
			return ones[n/100] + "hundred"
		} else {
			return ones[n/100] + "hundred" + "and" + Num2Words(n%100)
		}
	} else if n < 100000 {
		if n%1000 == 0 {
			return ones[n/1000] + "thousand"
		} else {
			return ones[n/1000] + "thousand" + Num2Words(n%1000)
		}
	}

	return ""
}

func main() {
	total := 0

	for i := 1; i <= 1000; i++ {
		total += len(Num2Words(i))
	}

	fmt.Println(total)
}
