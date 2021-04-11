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
	"strings"
)

type Fraction struct {
	n, d int
}

func (this *Fraction) Simplified() *Fraction {
	GreatestCommonDivisor := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}

		return a
	}

	gcd := GreatestCommonDivisor(this.n, this.d)

	return &Fraction{
		n: this.n / gcd,
		d: this.d / gcd,
	}
}

func (this *Fraction) Equals(f Fraction) bool {
	if *this.Simplified() == *f.Simplified() {
		return true
	}

	return false
}

func (this *Fraction) IsCurious() bool {
	n := strings.Split(strconv.Itoa(this.n), "")
	d := strings.Split(strconv.Itoa(this.d), "")

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if n[i] == d[j] {
				a, _ := strconv.Atoi(n[i^1])
				b, _ := strconv.Atoi(d[j^1])

				return this.Equals(Fraction{a, b})
			}
		}
	}

	return false
}

func main() {
	ns := 1
	ds := 1

	for i := 10; i < 100; i++ {
		for j := 10; j < 100; j++ {
			if !(i%10 == 0 || j%10 == 0 || i >= j) {
				f := Fraction{i, j}

				if f.IsCurious() {
					ns *= f.n
					ds *= f.d
				}
			}
		}
	}

	f := Fraction{ns, ds}
	f = *f.Simplified()
	fmt.Println(f.d)
}
