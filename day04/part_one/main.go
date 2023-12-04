package main

import (
    "os"
    "fmt"
    "math"
    "bufio"
    "unicode"
    "strings"
    "strconv"
)

func    convertStringToNum(s []string) []int {
    ret := make([]int, len(s))
    for i, val := range s {
        conv, _ := strconv.Atoi(val)
        ret[i] = conv
    }
    return ret
}

func    wsSplit(r rune) bool {
    return unicode.IsSpace(r)
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
    scanner := bufio.NewScanner(f)

    p1 := 0
    p2 := 0
    for scanner.Scan() {
        line := scanner.Text()
        content := strings.Split(line, ":")
        list := strings.Split(content[1], "|") 
        winingList := strings.TrimLeft(list[0], " ")
        toCmpList := strings.TrimLeft(list[1], " ")
        toCmpNum := convertStringToNum(strings.FieldsFunc(toCmpList, wsSplit))
        winingNum := convertStringToNum(strings.FieldsFunc(winingList, wsSplit))
        set := map[int]bool{}
        for _, x := range winingNum {
            for _, y := range toCmpNum {
                if x == y {
                    set[x] = true
                }
            }
        }
        win := len(set)
        p1 += int(math.Pow(2, float64(win-1)))
    }
    fmt.Println(p1)
    fmt.Println(p2)
}
