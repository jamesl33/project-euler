package main

import (
    "fmt"
    "math"
)

func RemoveDuplicates(slice []int) []int {
    known := map[int]bool {}
    result := []int {}

    for _, value := range slice {
        if !known[value] {
            known[value] = true
            result = append(result, value)
        }
    }

    return result
}

func Divisors(n int) []int {
    var divisors []int
    limit := int(math.Sqrt(float64(n)))

    for i := 1; i < limit + 1; i++ {
        if n % i == 0 {
            divisors = append(divisors, i, n / i)
        }
    }

    for index, item := range divisors {
        if item == n {
            divisors = append(divisors[:index], divisors[index + 1:]...)
            break
        }
    }

    return RemoveDuplicates(divisors)
}

func Sum(s []int) int {
    total := 0

    for _, item := range s {
        total += item
    }

    return total
}

func GenAbundantNumbers() []int {
    var abundantNumbers []int

    for i := 0; i < 28123; i++ {
        if Sum(Divisors(i)) > i {
            abundantNumbers = append(abundantNumbers, i)
        }
    }

    return abundantNumbers
}

func main() {
    abundantNumbers := GenAbundantNumbers()
    var expressible [28123]bool

    for _, i := range abundantNumbers {
        for _, j := range abundantNumbers {
            if i + j < 28123 {
                expressible[i + j] = true
            } else {
                break
            }
        }
    }

    total := 0

    for index, value := range expressible {
        if !value {
            total += index
        }
    }

    fmt.Println(total)
}
