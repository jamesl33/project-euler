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

// Mathimatical solution using Binomial coefficient. https://en.wikipedia.org/wiki/Binomial_coefficient
// With help from https://github.com/nayuki/Project-Euler-solutions/blob/master/python/p015.py

// package main

// import (
//     "fmt"
//     "math/big"
// )

// func factorial(n int64) *big.Int {
//     var result big.Int
//     result.MulRange(1, n)
//     return &result
// }

// func binomial(n int64, k int64) *big.Int {
//     if n < k || k < 0 {
//         panic("Error: Invalid input")
//     }

//     var result big.Int
//     var top big.Int
//     var bottom big.Int

//     top = *factorial(n)
//     bottom.Mul(factorial(k), factorial(n - k))
//     result.Div(&top, &bottom)

//     return &result
// }

// func main() {
//     fmt.Println(binomial(40, 20))
// }

// Programatic solution
// With help from https://euler.stephan-brumme.com/15/

package main

import "fmt"

type Coord struct {
	x, y int
}

func main() {
	gridSize := 20
	grid := make([][]int, gridSize+1)

	for index := range grid {
		grid[index] = make([]int, gridSize+1)
	}

	grid[gridSize][gridSize] = 1

	queue := make([]Coord, 0)
	queue = append(queue, Coord{gridSize - 1, gridSize})
	queue = append(queue, Coord{gridSize, gridSize - 1})

	for len(queue) != 0 {
		var current Coord

		current, queue = queue[0], queue[1:]

		if grid[current.x][current.y] == 0 {
			numRoutes := 0

			if current.x < gridSize {
				numRoutes += grid[current.x+1][current.y]
			}

			if current.y < gridSize {
				numRoutes += grid[current.x][current.y+1]
			}

			grid[current.x][current.y] = numRoutes

			if current.x > 0 {
				queue = append(queue, Coord{current.x - 1, current.y})
			}

			if current.y > 0 {
				queue = append(queue, Coord{current.x, current.y - 1})
			}
		}
	}

	fmt.Println(grid[0][0])
}
