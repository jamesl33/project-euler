package main

import "fmt"

func main() {
	for a := 1; a < 500; a++ {
		for b := 1; b < 500; b++ {
			for c := 1; c < 500; c++ {
				if (a*a)+(b*b) == (c*c) && a+b+c == 1000 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
