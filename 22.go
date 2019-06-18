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
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func LoadNames(fileName string) []string {
    var names []string

    file, _ := os.Open(fileName)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        names = append(names, strings.Split(scanner.Text(), ",")...)
    }

    return names
}

func main() {
    totalScore := 0
    names := LoadNames("names.txt")
    sort.Slice(names, func(i, j int) bool {return names[i] < names[j]})

    for index, name := range names {
        nameScore := 0

        for _, char := range name {
            if char != rune(34) {
                nameScore += int(char - 64)
            }
        }

        nameScore *= (index + 1)
        totalScore += nameScore
    }

    fmt.Println(totalScore)
}
