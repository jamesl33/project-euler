package main

import (
    "fmt"
    "math"
)

type Tuple struct {
    a, b int
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

    return divisors
}

func Sum(s []int) int {
    total := 0

    for _, item := range s {
        total += item
    }

    return total
}

func GenAmicablePairs(n int) []Tuple {
    var amicablePairs []Tuple

    for x := 1; x <= n; x++ {
        y := Sum(Divisors(x))

        if y > x && x == Sum(Divisors(y)) {
            amicablePairs = append(amicablePairs, Tuple{x, y})
        }
    }

    return amicablePairs
}

func main() {
    total := 0

    for _, tuple := range GenAmicablePairs(10000) {
        total += tuple.a
        total += tuple.b
    }

    fmt.Println(total)
}
