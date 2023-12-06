package main

import (
    "os"
    "fmt"
    "io"
    "strings"
    "math"
    "strconv"
)

type Race struct {
    time, dist int
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

    buf, _ := io.ReadAll(f)
    lines := strings.Split(string(buf), "\n")
    p2 := 0
    totalTime := ""
    totalDist := ""
    for i, line := range lines {
        vals := strings.Fields(strings.TrimLeft(line[strings.Index(line, ":")+1:], " \n\t"))
        for _, val := range vals {
            if i == 0 {
                totalTime += val 
            } else {
                totalDist += val
            }
        }
    }
    time, _ := strconv.Atoi(totalTime)
    dist, _ := strconv.Atoi(totalDist)
    realRace := Race{time, dist}
    for i := 0; i <= realRace.time; i++ {
        cal := i*int(math.Abs(float64(i - realRace.time)))
        if cal > realRace.dist {
            p2 += 1
        }
    }
    fmt.Println(p2)
}
