#include <iostream>

int fibonacci(int n) {
    if (n == 0 || n == 1) {
        return 1;
    } else {
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}

int main() {
    int currentNum = 0;
    int currentFib = 0;
    int total = 0;

    while (currentFib < 4000000) {
        currentFib = fibonacci(currentNum);

        if (currentFib % 2 == 0) {
            total += currentFib;
        }

        currentNum++;
    }

    std::cout << total << std::endl;
}
