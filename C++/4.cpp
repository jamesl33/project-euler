#include <iostream>

bool is_palindrome(int n) {
    int original = n;
    int reversed = 0;
    int remainder = 0;

    while (n != 0) {
        remainder = n % 10;
        reversed = reversed * 10 + remainder;
        n /= 10;
    }

    return original == reversed;
}

int main() {
    int current;
    int largest = 0;

    for (int i = 0; i < 1000; i++) {
        for (int j = 0; j < 1000; j++) {
            current = i * j;

            if (is_palindrome(current) && current > largest) {
                largest = current;
            }
        }
    }

    std::cout << largest << std::endl;
}
