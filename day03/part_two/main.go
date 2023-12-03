package main

import (
    "os"
    "io"
    "fmt"
    "strconv"
    "strings"
)

type Pos struct {
    y, x int
}

func    manAtoi(s string, i int) (int, error) {
    left, right := i, i
    for 0 <= left-1 && (s[left-1] >= '0' && s[left-1] <= '9') {
        left--
    }
    for right+1 < len(s) && (s[right+1] >= '0' && s[right+1] <= '9') {
        right++
    }
    return strconv.Atoi(s[left:right+1])
}

func    main() {
    input, err := os.Open("input")
    if err != nil {
        fmt.Println("error")
        os.Exit(1)
    }
    defer input.Close()

    p1 := 0
    p2 := 0
    raw, _ := io.ReadAll(input)
    file := strings.Split(string(raw), "\n")
    pos := []Pos{
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, -1},
        {0, 0},
        {0, 1},
        {1, -1},
        {1, 0},
        {1, 1},
    }
    file_len, line_len := len(file), len(file[0])
    for i, line := range file {
        for j, sym := range line {
            if strings.ContainsRune(".0123456789", sym) {
                continue
            }
            m := map[int]bool{}
            mult := map[rune][]int{}
            for _, idx := range pos {
                idx.y, idx.x =  idx.y + i, idx.x + j
                if ((idx.y >= 0 && idx.y <= file_len) && (idx.x >= 0 && idx.x <= line_len) && (file[idx.y][idx.x] >= '0' &&  file[idx.y][idx.x] <= '9')) {
                    num, _ := manAtoi(file[idx.y], idx.x)
                    if (!m[num]) {
                        m[num] = true
                        mult[sym] = append(mult[sym], num)
                    }
                }
            }
            if (len(mult['*']) == 2) {
                p2 += mult['*'][0] * mult['*'][1]
            }
            for k := range m {
                p1 += k
            }
        }
    }
    fmt.Println(p1)
    fmt.Println(p2)
}
