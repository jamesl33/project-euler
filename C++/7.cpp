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
    int currentNum = 0;
    int currentPrime = 0;
    int currentPrimeNum = 0;

    while (currentPrimeNum != 10001) {
        if (isPrime(currentNum)) {
            currentPrime = currentNum;
            currentPrimeNum++;
        }

        currentNum++;
    }

    std::cout << currentPrime << std::endl;
}
