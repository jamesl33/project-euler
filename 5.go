package main

import "fmt"

func is_divisible(n int) bool {
    for i := 1; i < 20; i++ {
        if !(n % i == 0) {
            return false
        }
    }
    return true
}

func main() {
    current_num := 1

    for true {
        if is_divisible(current_num) {
            fmt.Println(current_num)
            return
        }
        current_num++
    }
}
