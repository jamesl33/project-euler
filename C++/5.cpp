#include <iostream>

bool evenlyDivisible(int n) {
    for (int i = 1; i <= 20; i++) {
        if (!(n % i == 0)) {
            return false;
        }
    }

    return true;
}

int main() {
    int currentNum = 1;

    while (true) {
        if (evenlyDivisible(currentNum)) {
            std::cout << currentNum << std::endl;
            return 0;
        }

        currentNum++;
    }
}
