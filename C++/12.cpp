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
#include <cmath>

int Divisors(int n)
{
    int divisors = 0;

    for (int i = 1; i <= sqrt(n); i++)
    {
        if (n % i == 0)
        {
            if (n / i == 0)
            {
                divisors++;
            }
            else
            {
                divisors += 2;
            }
        }
    }

    return divisors;
}

int Triangle(int n)
{
    if (n == 1)
    {
        return n;
    }
    else
    {
        return n + Triangle(n - 1);
    }
}

int main()
{
    int current = 1;

    while (Divisors(Triangle(current)) < 500)
    {
        current++;
    }

    std::cout << Triangle(current) << std::endl;
}
