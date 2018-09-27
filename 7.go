package main

import "fmt"

func is_prime(n int) bool {
    if n <= 1 {
        return false
    } else if n <= 3 {
        return true
    } else if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    i := 5

    for i * i <= n {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }

        i += 6
    }

    return true
}

func main() {
    primes_found := 0
    last_prime := 0
    current_num := 0

    for primes_found <= 10000 {
        if is_prime(current_num) {
            primes_found++
            last_prime = current_num
        }

        current_num++
    }

    fmt.Println(last_prime)
}
