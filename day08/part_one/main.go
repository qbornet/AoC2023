package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Node struct {
    left, right string
}

type Finished struct {
    step int
    finished bool
}


func    ContainsAtEnd(s string) bool {
    return s[len(s)-1] == 'Z'
}

func    GCD(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func    LCM(a, b int, ints ...int) int {
    result := a*b / GCD(a, b)

    for _, i := range ints {
        result = LCM(result, i)
    }
    return result
}

func    main() {
    f, err := os.Open("input")
    if err != nil {
        f = os.Stdin
    }

    if len(os.Args) == 2 {
        f.Close()
        f, _ = os.Open(os.Args[1])
    }
    defer f.Close()

    p1 := 0
    buf, _ := io.ReadAll(f)
    lines := strings.Split(string(buf), "\n")
    mouv := lines[0]
    lines = lines[2:len(lines)-1]
    net := make(map[string]Node)
    for _, line := range lines {
        content := strings.Split(line, "=")
        pos := strings.TrimRight(content[0], " ")
        all := strings.TrimLeft(content[1], " ")
        all = all[1:len(all)-1]
        dst := strings.Split(all, ", ")
        net[pos] = Node{dst[0], dst[1]}
    }
    currentPos := "AAA"
    for currentPos != "ZZZ" {
        pos := net[currentPos]
        for _, c := range mouv {
            p1++
            if c == 'R' {
                currentPos = pos.right
            } else {
                currentPos = pos.left
            }
            if (currentPos == "ZZZ") {
                break
            }
            pos = net[currentPos]
        }
    }
    fmt.Println(p1)
}
