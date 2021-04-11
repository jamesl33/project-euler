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
	"math/big"
)

func FibSequence(n int) []*big.Int {
	sequence := []*big.Int{big.NewInt(0), big.NewInt(1)}

	for i := 2; i < n; i++ {
		var nextItem big.Int

		sequence = append(sequence, nextItem.Add(sequence[i-1], sequence[i-2]))
	}

	return sequence
}

func main() {
	current := 0

	for true {
		sequence := FibSequence(current)

		if len(sequence[len(sequence)-1].String()) == 1000 {
			fmt.Println(len(sequence) - 1)
			break
		}

		current++
	}
}
