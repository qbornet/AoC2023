package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

const (
    max_red = iota + 12
    max_green
    max_blue
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
        red, green, blue := false, false, false
        line := scanner.Text()
        gameId := line[5:strings.Index(line, ":")]
        setCube := strings.Split(line[strings.Index(line, ":")+1:], ";")
        for _, set := range setCube {
            ret := strings.Split(set, ",")
            for _, cube := range ret {
                cube = strings.TrimLeft(cube, " ")
                val := strings.Index(cube, "red")
                if (val > 0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (conv > max_red) {
                        red = true
                    }
                    continue
                }
                val = strings.Index(cube, "green")
                if (val >  0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (conv > max_green) {
                        green = true
                    }
                    continue
                }
                val = strings.Index(cube, "blue")
                if (val >  0) {
                    conv, _ := strconv.Atoi(cube[:strings.Index(cube, " ")])
                    if (conv > max_blue) {
                        blue = true
                    }
                    continue
                }
            }
        }
        if (red == true || green == true || blue == true) {
            continue
        }
        conv, _ := strconv.Atoi(gameId)
        r += conv
    }
    fmt.Println(r)
}
