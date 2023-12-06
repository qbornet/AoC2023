package main

import (
	"bufio"
	"fmt"
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

	var seeds [][2]int
	tables := make([][][3]int, 0)
	var table [][3]int
	prev := lineKind(-1)

	saveTable := func() {
		slices.SortFunc(table, func(a, b [3]int) int {
			return a[0] - b[0]
		})
		tables = append(tables, table)
	}

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		if prev == Empty {
			prev = Header
			continue
		}
		line := strings.Split(scanner.Text(), " ")
		if len(line) == 1 {
			prev = Empty
			continue
		}
		if row == 0 {
			pairs := make([]int, len(line)-1)
			setInts(line[1:], pairs)
			seeds = make([][2]int, len(pairs)/2)
			for i := 0; i < len(seeds); i++ {
				seeds[i] = [...]int{pairs[i*2], pairs[i*2+1]}
			}
			slices.SortFunc(seeds, func(a, b [2]int) int {
				return a[0] - b[0]
			})
			prev = Seeds
			continue
		}
		if prev == Header {
			if table != nil {
				saveTable()
			}
			table = make([][3]int, 0)
		}
		if prev == Header || prev == Mapping {
			var entry [3]int
			setInts(line, entry[:])
			table = append(table, entry)
			prev = Mapping
			continue
		}
	}
	saveTable()

	slices.Reverse(tables)

all:
	for l := 0; true; l++ {
		m := l
		for _, table := range tables {
			for _, e := range table {
				a := e[0]
				if a <= m && m <= a+e[2]-1 {
					m = m - a + e[1]
					break
				}
			}
		}

		for _, seed := range seeds {
			start := seed[0]
			if m < start {
				continue all
			}
			if m <= start+seed[1]-1 {
				fmt.Println(l)
				break all
			}
		}
	}
}
