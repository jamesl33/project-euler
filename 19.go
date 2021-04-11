package main

import "fmt"

func isLeapYear(year int) bool {
	return year%4 == 0 && !(year%100 == 0 && year%400 != 0)
}

func main() {
	var (
		currentDay, totalSundays int
		months                   = []int{31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	)

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
