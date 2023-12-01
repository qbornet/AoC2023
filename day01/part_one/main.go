package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
)

func    main() {
    input, err := os.Open("input")
    defer input.Close()
    if err != nil {
        fmt.Println("error")
        os.Exit(1)
    }
    scanner := bufio.NewScanner(input)

    total := 0
    for scanner.Scan() {
        first := 0
        line := scanner.Text()
        num := make([]rune, 2)
        num[1] = 0
        for _, c := range line {
            if first == 0 && c >= '0' && c <= '9' {
                num[0] = c
                first = 1
            } else if c >= '0' && c <= '9' {
                num[1] = c
            }
        }
        if num[1] == 0 {
            num[1] = num[0]
        }
        toAdd, _ := strconv.Atoi(string(num))
        total += toAdd
    }
    fmt.Println(total)
}
