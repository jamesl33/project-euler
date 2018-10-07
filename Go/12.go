package main

import (
    "fmt"
    "math"
)

func divisors(n int) []int {
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

    return divisors
}

func triangle(n int) int {
    if n == 1 {
        return 1
    } else {
        return n + triangle(n - 1)
    }
}

func main() {
    current := 1

    for true {
        triangle := triangle(current)

        if len(divisors(triangle)) > 500 {
            fmt.Println(triangle)
            return
        }

        current++
    }
}
