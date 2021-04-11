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

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}

		return true
	}

	return false
}

func main() {
	currentDay := 0
	totalSundays := 0
	months := []int{31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for year := 1901; year <= 2000; year++ {
		for _, monthDays := range months {
			if monthDays == 0 {
				if isLeapYear(year) {
					monthDays = 29
				} else {
					monthDays = 28
				}
			}

			if currentDay%7 == 0 {
				totalSundays++
			}

			currentDay += monthDays
		}
	}

	fmt.Println(totalSundays)
}
