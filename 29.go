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

package main

import (
    "fmt"
    "math/big"
)

func contains(slice []big.Int, n *big.Int) bool {
    for _, e := range slice {
        if e.Cmp(n) == 0 {
            return true
        }
    }

    return false
}

func main() {
    var powers []big.Int

    for a := 2; a <= 100; a++ {
        for b := 2; b <= 100; b++ {
            var result big.Int
            result.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
            powers = append(powers, result)
        }
    }

    var unique_powers []big.Int

    for _, p := range powers {
        if !contains(unique_powers, &p) {
            unique_powers = append(unique_powers, p)
        }
    }

    fmt.Println(len(unique_powers))
}
