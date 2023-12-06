package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type lineKind int

const (
	Seeds   lineKind = 0
	Empty   lineKind = 1
	Header  lineKind = 2
	Mapping lineKind = 3
)

func setInts(src []string, tgt []int) {
	for i, str := range src {
		seed, _ := strconv.Atoi(str)
		tgt[i] = seed
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
        file = os.Stdin
	}

    if len(os.Args) == 2 {
        file.Close()
        file, _ = os.Open(os.Args[1])
    }
	defer file.Close()

	var seeds []int
	var mapped []int
	var done []bool
	var mapping [3]int

	prev := lineKind(-1)

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		if prev == Empty {
			prev = Header
			continue
		}
		line := strings.Split(scanner.Text(), " ")
		if len(line) == 1 {
			for i := range done {
				done[i] = false
			}
			prev = Empty
			continue
		}
		if row == 0 {
			seeds = make([]int, len(line)-1)
			setInts(line[1:], seeds)
			mapped = slices.Clone(seeds)
			done = make([]bool, len(seeds))
			prev = Seeds
			continue
		}
		if prev == Header || prev == Mapping {
			setInts(line, mapping[:])
			a := mapping[0]
			b := mapping[1]
			c := mapping[2]
			for i, d := range done {
				if d {
					continue
				}
				if x := mapped[i]; b <= x && x <= b+c-1 {
					mapped[i] = a + x - b
					done[i] = true
				}
			}
			prev = Mapping
			continue
		}
	}
	lowest := math.MaxInt
	for _, x := range mapped {
		lowest = min(x, lowest)
	}
    fmt.Println(lowest)
}
