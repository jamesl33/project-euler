#include <iostream>

int main() {
    int squareSum = 0;
    int sum = 0;

    for (int i = 1; i <= 100; i++) {
        squareSum += i * i;
        sum += i;
    }

    std::cout << (sum * sum) - squareSum << std::endl;
}
