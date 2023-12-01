package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
)
const (
    one int = iota + 1
    two 
    three
    four
    five
    six
    seven
    eight
    nine
)

func    checkStringIsNumber(line, s string, current, endPos, lineLength int) int {
    switch s {
    case "one":
        return one
    case "two":
        return two
    case "three":
        return three
    case "four":
        return four
    case "five":
        return five
    case "six":
        return six
    case "seven":
        return seven
    case "eight":
        return eight
    case "nine":
        return nine
    default:
        if ((endPos - current) > len("eight")) {
            return 0
        }
        if (endPos <= lineLength) {
            return checkStringIsNumber(line, line[current:endPos], current, endPos+1, lineLength)
        }
        return 0
    }
}

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
        for i, c := range line {
            val := 0
            lineLength := len(line)
            if (c == 'o' || c == 't' || c == 'f' || c == 's' || c == 'e' || c == 'n') {
                val = checkStringIsNumber(line, line, i, i+1, lineLength)
            }
            if first == 0 && val > 0 {
                num[0] = rune(strconv.Itoa(val)[0])
                first = 1
            } else if val > 0 {
                num[1] = rune(strconv.Itoa(val)[0])
            }

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
