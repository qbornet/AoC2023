package main

import (
    "os"
    "io"
    "fmt"
    "strings"
    "strconv"

    "github.com/atotto/clipboard"
)

func    equalZero(a []int) bool {
    for _, val := range a {
        if val != 0 {
            return false
        }
    }
    return true
}

func    findNext(arr []int) int {
    if (equalZero(arr)) {
        return 0
    }

    x := arr[0]
    iter := arr[1:]
    diff := []int{}
    for _, y := range iter {
        diff = append(diff, y - x)
        x = y
    }
    res := findNext(diff)
    return arr[0] - res
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

    p2 := 0
    buf, _ := io.ReadAll(f)
    lines := strings.Split(string(buf), "\n")
    lines = lines[:len(lines)-1]
    for _, line := range lines {
        nums := strings.Fields(line)
        arr := []int{}
        for _, num := range nums {
            conv, _ := strconv.Atoi(num)
            arr = append(arr, conv)
        }
        p2 += findNext(arr)
    }
    clipboard.WriteAll(strconv.Itoa(p2))
    fmt.Println(p2)
}
