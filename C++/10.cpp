#include <iostream>

bool isPrime(int n) {
    if (n <= 1) {
        return false;
    } else if (n <= 3) {
        return true;
    } else if (n % 2 == 0 || n % 3 == 0) {
        return false;
    }

    int i = 5;

    while (i * i <= n) {
        if (n % i == 0 || n % (i + 2) == 0) {
            return false;
        }

        i += 6;
    }

    return true;
}

int main() {
    long int total = 0;

    for (int i = 0; i < 2000000; i++) {
        if (isPrime(i)) {
            total += i;
        }
    }

    std::cout << total << std::endl;
}
