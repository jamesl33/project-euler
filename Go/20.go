package main

import (
    "fmt"
    "math/big"
)

func factorial(n int64) *big.Int {
    var result big.Int
    result.MulRange(1, n)
    return &result
}

func main() {
    total := 0

    for _, char := range factorial(100).String() {
        total += int(char - '0')
    }

    fmt.Println(total)
}
