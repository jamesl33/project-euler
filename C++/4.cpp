/* This file is part of project-euler

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

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
