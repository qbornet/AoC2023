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

    p1 := 1
    buf, _ := io.ReadAll(f)
    lines := strings.Split(string(buf), "\n")
    var races []Race
    for i, line := range lines {
        vals := strings.Fields(strings.TrimLeft(line[strings.Index(line, ":")+1:], " \n\t"))
        for j, val := range vals {
            conv, _ := strconv.Atoi(val)
            if i == 0 {
                races = append(races, Race{conv, 0})
            } else {
                races[j] = Race{races[j].time, conv}
            }
        }
    }
    for _, race := range races {
        total := 0
        for i := 0; i <= race.time; i++ {
            cal := i*int(math.Abs(float64(i - race.time)))
            if cal > race.dist {
                total += 1
            }
        }
        p1 *= total
    }
    fmt.Println(p1)
}
