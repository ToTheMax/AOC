package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id            int
	items         []int
	inspected     int
	operation     string
	divisibleTest int
	throwMonkey1  int
	throwMonkey2  int
}

func parseMonkeys(input string) []*Monkey {
	monkeyInputs := strings.Split(input, "\n\n")
	monkeys := make([]*Monkey, len(monkeyInputs))
	for _, monkeyInput := range monkeyInputs {
		monkeyInput = strings.ReplaceAll(monkeyInput, " ", "")
		inputLines := strings.Split(monkeyInput, "\n")
		monkey := &Monkey{}
		itemsString := ""
		fmt.Sscanf(inputLines[0], "Monkey%d:", &monkey.id)
		fmt.Sscanf(inputLines[1], "Startingitems:%s", &itemsString)
		fmt.Sscanf(inputLines[2], "Operation:new=old%s", &monkey.operation)
		fmt.Sscanf(inputLines[3], "Test:divisibleby%d", &monkey.divisibleTest)
		fmt.Sscanf(inputLines[4], "Iftrue:throwtomonkey%d", &monkey.throwMonkey1)
		fmt.Sscanf(inputLines[5], "Iffalse:throwtomonkey%d", &monkey.throwMonkey2)
		splitted_items := strings.Split(itemsString, ",")
		items := make([]int, len(splitted_items))
		for i, itemStr := range splitted_items {
			items[i], _ = strconv.Atoi(itemStr)
		}
		monkey.items = items
		monkeys[monkey.id] = monkey
	}
	return monkeys
}

func expression(old int, expression string) int {
	if expression == "old" {
		return old
	} else {
		value, _ := strconv.Atoi(expression)
		return value
	}
}

func operation(old int, operation string) int {
	switch operation[0] {
	case '*':
		return old * expression(old, operation[1:])
	case '+':
		return old + expression(old, operation[1:])
	case '-':
		return old - expression(old, operation[1:])
	default:
		return 0
	}
}

func inspection(monkey *Monkey, monkeys []*Monkey, worryDivision int, mod int) {
	for _, item := range monkey.items[monkey.inspected:] {
		monkey.inspected++

		// Multiply worry level
		new := operation(item, monkey.operation)
		new = new % mod
		new = new / worryDivision

		// Throw item
		nextMonkey := monkeys[monkey.throwMonkey2]
		if new%monkey.divisibleTest == 0 {
			nextMonkey = monkeys[monkey.throwMonkey1]
		}

		nextMonkey.items = append(nextMonkey.items, new)
	}

}

func inspectionCounts(monkeys []*Monkey) []int {
	inspectionCounts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspectionCounts[i] = monkey.inspected
	}
	return inspectionCounts
}

func startRounds(rounds int, worryDivision int, monkeys []*Monkey) int {
	mod := 1
	for _, monkey := range monkeys {
		mod *= monkey.divisibleTest
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			inspection(monkey, monkeys, worryDivision, mod)
		}
	}
	inspectionCounts := inspectionCounts(monkeys)
	sort.Sort(sort.Reverse(sort.IntSlice(inspectionCounts)))
	return inspectionCounts[0] * inspectionCounts[1]
}

func main() {
	input, _ := os.ReadFile("in.txt")

	fmt.Println("Sol1:", startRounds(20, 3, parseMonkeys(string(input))))
	fmt.Println("Sol2:", startRounds(10000, 1, parseMonkeys(string(input))))
}
