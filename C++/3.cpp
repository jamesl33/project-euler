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
