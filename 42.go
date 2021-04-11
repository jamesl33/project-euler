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
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadWords(filename string) []string {
	var words []string

	if file, err := os.Open(filename); err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			words = append(words, strings.Split(scanner.Text(), ",")...)
		}
	}

	return words
}

func IsTriangleNum(n int) bool {
	nthTriangle := (math.Sqrt(float64(1+8*n)) - 1) / 2

	if nthTriangle == float64(int(nthTriangle)) {
		return true
	}

	return false
}

func IsTriangleWord(s string) bool {
	sum := 0

	for _, c := range s {
		sum += int(c) - 96
	}

	return IsTriangleNum(sum)
}

func main() {
	count := 0

	for _, word := range ReadWords("p042_words.txt") {
		word = strings.ToLower(word)
		word = strings.Trim(word, "\"")

		if IsTriangleWord(word) {
			count++
		}
	}

	fmt.Println(count)
}
