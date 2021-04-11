package main

import (
	"fmt"
	"strconv"
	"strings"
)

type fraction struct {
	n, d int
}

func (f *fraction) Simplified() *fraction {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}

		return a
	}

	d := gcd(f.n, f.d)

	return &fraction{
		n: f.n / d,
		d: f.d / d,
	}
}

func (f *fraction) Equals(other fraction) bool {
	if *f.Simplified() == *other.Simplified() {
		return true
	}

	return false
}

func (f *fraction) IsCurious() bool {
	n := strings.Split(strconv.Itoa(f.n), "")
	d := strings.Split(strconv.Itoa(f.d), "")

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if n[i] == d[j] {
				a, _ := strconv.Atoi(n[i^1])
				b, _ := strconv.Atoi(d[j^1])

				return f.Equals(fraction{a, b})
			}
		}
	}

	return false
}

func main() {
	ns, ds := 1, 1

	for i := 10; i < 100; i++ {
		for j := 10; j < 100; j++ {
			if !(i%10 == 0 || j%10 == 0 || i >= j) {
				f := fraction{i, j}

				if f.IsCurious() {
					ns *= f.n
					ds *= f.d
				}
			}
		}
	}

	fmt.Println((&fraction{ns, ds}).Simplified().d)
}
