package main

import "fmt"

func LargestPrimeFactor(n int) int {
    current_largest := 0
    divisor := 2

    for n > 1 {
        for n % divisor == 0 {
            current_largest = divisor
            n /= divisor
        }

        divisor++
    }

    return current_largest
}

func main() {
    fmt.Println(LargestPrimeFactor(600851475143))
}
