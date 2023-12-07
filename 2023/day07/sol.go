package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	name  string
	score int
	count int
}

type Hand struct {
	rank  int
	score int
	bid   int
}

func cardScore(char rune) int {
	ascii_num := int(char - '0')
	if ascii_num >= 0 && ascii_num <= 9 {
		return ascii_num
	} else {
		if char == 'T' {
			return 10
		} else if char == 'J' {
			return 11
		} else if char == 'Q' {
			return 12
		} else if char == 'K' {
			return 13
		} else if char == 'A' {
			return 14
		}
	}
	return -1
}

func handScore(hand string, bid int) Hand {
	cardsMap := map[string]*Card{}
	for _, char := range hand {
		name := string(char)
		if _, ok := cardsMap[name]; ok {
			cardsMap[name].count++
		} else {
			cardsMap[name] = &Card{
				name:  string(char),
				score: cardScore(char),
				count: 1,
			}
		}
	}

	cards := make([]Card, 0)
	for _, card := range cardsMap {
		cards = append(cards, *card)
	}

	sort.Slice(cards[:], func(i, j int) bool {
		// First compare counts, then score
		cardi := cards[i].count*100 + cards[i].score
		cardj := cards[j].count*100 + cards[j].score
		return cardi > cardj
	})

	// Check for Five of a kind
	if cards[0].count == 5 {
		return Hand{6, cards[0].score, bid}
	}
	// Check for Four of a kind
	if cards[0].count == 4 {
		return Hand{5, 10*cards[0].score + cards[1].score, bid}
	}

	// Check for Full House
	if cards[0].count == 3 && cards[1].count == 2 {
		return Hand{4, 10*cards[0].score + cards[1].score, bid}
	}

	// Check for Three of a kind
	if cards[0].count == 3 {
		return Hand{3, 100*cards[0].score + 10*cards[1].score + cards[2].score, bid}
	}

	// Check for Two Pairs
	if cards[0].count == 2 && cards[1].count == 2 {
		return Hand{2, 100*cards[0].score + 10*cards[1].score + cards[2].score, bid}
	}

	// Check for Single Pair
	if cards[0].count == 2 {
		return Hand{1, 1000*cards[0].score + 100*cards[1].score + 10*cards[2].score + cards[3].score, bid}
	}

	// Check for High Card
	return Hand{0, 10000*cards[0].score + 1000*cards[1].score + 100*cards[2].score + 10*cards[3].score + cards[4].score, bid}
}

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		splittedLine := strings.Split(line, " ")
		hand := splittedLine[0]
		bid, _ := strconv.Atoi(splittedLine[1])
		hands[i] = handScore(hand, bid)
		// fmt.Println(handScore(hand, bid))
	}

	sort.Slice(hands[:], func(i, j int) bool {
		// First compare ranks, then score
		cardi := hands[i].rank*100000 + hands[i].score
		cardj := hands[j].rank*100000 + hands[j].score
		return cardi < cardj
	})

	totalWinnings := 0
	for i, hand := range hands {
		fmt.Println(i+1, "*", hand.bid)
		totalWinnings += (i + 1) * hand.bid

	}
	fmt.Println("Sol 1", totalWinnings)
}
