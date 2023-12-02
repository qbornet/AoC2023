package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func    main() {
    input, err := os.Open("input")
    defer input.Close()
    if err != nil {
        os.Exit(1)
    }

    scanner := bufio.NewScanner(input)
    r := 0
    for scanner.Scan() {
        contains := make(map[string]int)
        contains["red"] = 0
        contains["green"] = 0
        contains["blue"] = 0
        line := scanner.Text()
        setCube := strings.Split(line[strings.Index(line, ":")+1:], ";")
        for _, set := range setCube {
            ret := strings.Split(set, ",")
            for _, cube := range ret {
                cube = strings.TrimLeft(cube, " ")
                val := strings.Index(cube, "red")
                if (val > 0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (contains["red"] < conv) {
                        contains["red"] = conv
                    }
                    continue
                }
                val = strings.Index(cube, "green")
                if (val >  0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (contains["green"] < conv) {
                        contains["green"] = conv
                    }
                    continue
                }
                val = strings.Index(cube, "blue")
                if (val >  0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (contains["blue"] < conv) {
                        contains["blue"] = conv
                    }
                    continue
                }
            }
        }
        r += contains["red"] * contains["green"] * contains["blue"]
    }
    fmt.Println(r)
}
