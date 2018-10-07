package main

import "fmt"

func main() {
    var collatz = [][]int {}

    for i := 1; i <= 1000000; i++ {
        sequence := []int {}

        n := i

        for n != 1 {
            if n % 2 == 0 {
                    n /= 2
            } else {
                n = n * 3 + 1
            }

            sequence = append(sequence, n)
        }

        collatz = append(collatz, sequence)
    }

    longest := 0

    for index, sequence := range collatz {
        if len(sequence) > len(collatz[longest]) {
            longest = index
        }
    }

    fmt.Println(longest + 1)
}
