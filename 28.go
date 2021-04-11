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

// func main() {
//     // 3 13 31 = 4(n * n) - 2n + 1
//     // 5 17 37 = 4(n * n) + 1
//     // 7 21 43 = 4(n * n) + 2n + 1
//     // 9 25 49 = 4(n * n) + 4n + 1

//     sum := 0

//     for i := 1; i <= (1001 - 1) / 2; i++ {
//         sum += 4 * (i * i) - (2 * i) + 1
//         sum += 4 * (i * i) + 1
//         sum += 4 * (i * i) + (2 * i) + 1
//         sum += 4 * (i * i) + (4 * i) + 1
//     }

//     // add 1 because the sequences do not take into account the 1 in the centre
//     fmt.Println(sum + 1)
// }

func main() {
	sum := 0

	for i := 1; i*i <= 1001*1001; i += 2 {
		for w := 0; w < 4; w++ {
			sum += (i * i) - ((i - 1) * w)
		}
	}

	// minus 3 because centre '1' will be processed 4 times overall
	fmt.Println(sum - 3)
}
