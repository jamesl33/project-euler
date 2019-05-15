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
