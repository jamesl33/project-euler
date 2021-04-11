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

func main() {
	max, maxp := 0, 0

	for p := 2; p <= 1000; p += 2 {
		count := 0

		for a := 2; a < p/3; a++ {
			if p*(p-2*a)%(2*(p-a)) == 0 {
				count++
			}
		}

		if count > max {
			max = count
			maxp = p
		}
	}

	fmt.Println(maxp)
}
