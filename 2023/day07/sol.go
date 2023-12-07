package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	str string
	bid   int
	rank  int
}

func cardScore(char rune, level int) int {
	if level == 1{
		return strings.Index("23456789TJQKA", string(char))
	} else{
		return strings.Index("J23456789TQKA", string(char))
	}
}


func handRank(str string, bid int, level int) Hand {
	cardCounts := make([]int, 13)
	for _, char := range str {
		cardCounts[cardScore(char, level)]++
	}
	
	jokers := 0
	if level == 2 {
		// Replace Jokers with 0
		jokers = cardCounts[0]
		cardCounts[0] = 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cardCounts)))

	if level == 2 {
		cardCounts[0] += jokers
	}
	rankings := map[int][]int{
		7: {5,0}, // Five of a kind
		6: {4,1}, // Four of a kind
		5: {3,2}, // Full House
		4: {3,1}, // Three of a kind
		3: {2,2}, // Two Pairs
		2: {2,1}, // Single Pair
		1: {1,1}, // High Card
	}

	for rank, counts := range rankings{
		if cardCounts[0] == counts[0] && cardCounts[1] == counts[1]{
			return Hand{str, bid, rank}
		}
	}
	return Hand{str, bid, -1}
}

func sortHands(hands []Hand, level int) {
	sort.Slice(hands[:], func(i, j int) bool {
		// First compare ranks, then string
		if hands[i].rank != hands[j].rank {
			return hands[i].rank < hands[j].rank
		} else {
			for k := range hands[i].str {
				if hands[i].str[k] != hands[j].str[k] {
					cardi := cardScore(rune(hands[i].str[k]),level)
					cardj := cardScore(rune(hands[j].str[k]),level)
					return cardi < cardj	
				}
			}
		}
		return false
	})
}


func calcScore(hands []Hand, level int) int {
	sortHands(hands, level)
	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}
	return totalWinnings
}

func main() {

	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		splittedLine := strings.Split(line, " ")
		hand := splittedLine[0]
		bid, _ := strconv.Atoi(splittedLine[1])
		hands[i] = handRank(hand, bid, 1)
	}

	fmt.Println("Sol 1:", calcScore(hands, 1))
	
	for i, hand := range hands{
		hands[i] = handRank(hand.str, hand.bid, 2)
	}

	fmt.Println("Sol 2:", calcScore(hands, 2))
	
}
