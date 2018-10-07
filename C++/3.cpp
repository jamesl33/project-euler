#include <iostream>

int LargestPrimeFactor(long int n) {
    int divisor = 2;
    int largest = 0;

    while (n > 1) {
        while (n % divisor == 0) {
            largest = divisor;
            n /= divisor;
        }

        divisor++;
    }

    return largest;
}

int main() {
    std::cout << LargestPrimeFactor(600851475143) << std::endl;
}
