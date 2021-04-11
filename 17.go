package main

import "fmt"

func numberToWords(n int) string {
	ones := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if 0 <= n && n < 20 {
		return ones[n]
	} else if n < 100 {
		if n%10 == 0 {
			return tens[n/10]
		} else {
			return tens[n/10] + ones[n%10]
		}
	} else if n < 1000 {
		if n%100 == 0 {
			return ones[n/100] + "hundred"
		} else {
			return ones[n/100] + "hundred" + "and" + numberToWords(n%100)
		}
	} else if n < 100000 {
		if n%1000 == 0 {
			return ones[n/1000] + "thousand"
		} else {
			return ones[n/1000] + "thousand" + numberToWords(n%1000)
		}
	}

	return ""
}

func main() {
	total := 0

	for i := 1; i <= 1000; i++ {
		total += len(numberToWords(i))
	}

	fmt.Println(total)
}
