package main

import (
    "fmt"
    "math/big"
)

func main() {
    var result big.Int
    result.Exp(big.NewInt(2), big.NewInt(1000), nil)
    resString := result.String()

    total := 0

    for _, char := range resString {
        total += int(char - '0')
    }

    fmt.Println(total)
}
