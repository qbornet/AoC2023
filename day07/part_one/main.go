package main

import (
    "io"
    "os"
    "fmt"
    "slices"
    "strings"
    "strconv"
)

const (
    high = iota
    one
    two
    three
    full
    four
    five
)


type Hands struct {
    bid int
    print_hand string
    hand string
}

var (
    cards = map[rune]int{
        rune('2'): 1,
        rune('3'): 2,
        rune('4'): 3,
        rune('5'): 4,
        rune('6'): 5,
        rune('7'): 6,
        rune('8'): 7,
        rune('9'): 8,
        rune('9'+1): 9,
        rune('9'+2): 10,
        rune('9'+3): 11,
        rune('9'+4): 12,
        rune('9'+5): 13,
    }
    handsKind = make(map[int][]Hands)
)

func    checkHands(h1 Hands, h2 Hands) int {
    i := 0
    for ;i < len(h1.hand); i++ {
        if (h1.hand[i] != h2.hand[i]) {
            break
        }
    }
    if (i == len(h1.hand)) {
        return 0
    }
    return cards[rune(h1.hand[i])] - cards[rune(h2.hand[i])]
} 

func    placeKind(c *map[int]int, h *Hands) {
    counter := *c
    if (counter[1] == 5) {
        handsKind[high] = append(handsKind[high], *h)
    }
    if (counter[1] == 3 && counter[2] == 1) {
        handsKind[one] = append(handsKind[one], *h)
    }
    if (counter[1] == 1 && counter[2] == 2) {
        handsKind[two] = append(handsKind[two], *h)
    }
    if (counter[1] == 2 && counter[3] == 1) {
        handsKind[three] = append(handsKind[three], *h)
    }
    if (counter[1] == 1 && counter[4] == 1) {
        handsKind[four] = append(handsKind[four], *h)
    }
    if (counter[5] == 1) {
        handsKind[five] = append(handsKind[five], *h)
    }
    if (counter[2] == 1 && counter[3] == 1) {
        handsKind[full] = append(handsKind[full], *h)
    }
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
    lines = lines[:len(lines)-1]
    for _, line := range lines{
        var hand Hands

        res := strings.Split(line, " ")
        cards := res[0]
        b := res[1]
        bid, _ := strconv.Atoi(b)

        hand.print_hand = cards
        hand.bid = bid
        cards = strings.ReplaceAll(cards, "T", string('9'+1))
        cards = strings.ReplaceAll(cards, "J", string('9'+2))
        cards = strings.ReplaceAll(cards, "Q", string('9'+3))
        cards = strings.ReplaceAll(cards, "K", string('9'+4))
        cards = strings.ReplaceAll(cards, "A", string('9'+5))
        hand.hand = cards
        hashRune := make(map[rune]int)
        for _, c := range cards {
            hashRune[c] += 1
        }
        counter := make(map[int]int)
        for _, val := range hashRune {
            counter[val] += 1
        }
        placeKind(&counter, &hand)
    }
    count := 1
    for i := 0; i <= len(handsKind); i++ {
        hands := handsKind[i]
        slices.SortFunc(hands, checkHands)
        for _, h := range hands {
            p1 += h.bid * count
            count++
        }
    }
    fmt.Println(p1)
}
