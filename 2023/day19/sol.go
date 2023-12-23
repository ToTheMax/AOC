package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	field string
	less   bool
	value int
	dest  *WorkFlow
}

type WorkFlow struct {
	name  string
	rules []Rule
}

type Part = map[string]int
type PartRange struct {
	lower int 
	upper int
}
type PartRanges = map[string]PartRange

func parseRule(line string, workFlows map[string]*WorkFlow) {
	split1 := strings.SplitN(line, "{", 2)
	workFlowName := split1[0]
	rules := strings.Split(split1[1][:len(split1[1])-1], ",")
	for _, rule := range rules {
		split2 := strings.SplitN(rule, ":", 2)
		if len(split2) == 1 {
			workFlows[workFlowName].rules = append(
				workFlows[workFlowName].rules, Rule{"", false, 0, workFlows[split2[0]]},
			)
		} else {
			condition := split2[0]
			val, _ := strconv.Atoi(condition[2:])
			workFlows[workFlowName].rules = append(
				workFlows[workFlowName].rules, Rule{condition[:1], condition[1] == '<', val, workFlows[split2[1]]},
			)
		}
	}
}

func makePart(part Part, workFlow WorkFlow) int {
	// Return Result
	if workFlow.name == "A" {
		return part["x"] + part["m"] + part["a"] + part["s"]
	} else  if workFlow.name == "R" {
		return 0
	}
	// Check rules
	for _, rule := range workFlow.rules {
		if rule.field == "" {
			return makePart(part, *rule.dest)
		} else if rule.less {
			if part[rule.field] < rule.value {
				return makePart(part, *rule.dest)
			}
		} else {
			if part[rule.field] > rule.value {
				return makePart(part, *rule.dest)
			}
		}
	}
	return 0
}

func makePartRanges(ranges PartRanges, workFlow WorkFlow) int {
	if workFlow.name == "A" {
		product := 1
		for _, partRange := range ranges {
			product *= partRange.upper - partRange.lower + 1
		}
		return product
	} else if workFlow.name == "R" {
		return 0
	}

	total := 0
	for _, rule := range workFlow.rules {
		newRanges := PartRanges{"x": ranges["x"], "m": ranges["m"], "a": ranges["a"], "s": ranges["s"]}
		pr := ranges[rule.field]
		if rule.field == "" {
			total += makePartRanges(newRanges, *rule.dest)
		} else if rule.less {
			newRanges[rule.field] = PartRange{pr.lower, rule.value - 1}
			ranges[rule.field] = PartRange{rule.value, pr.upper}
			total += makePartRanges(newRanges, *rule.dest)
		} else {
			newRanges[rule.field] = PartRange{rule.value + 1, pr.upper}
			ranges[rule.field] = PartRange{pr.lower, rule.value}
			total += makePartRanges(newRanges, *rule.dest)
		}
	}
	return total
}
	

func main() {
	input, _ := os.ReadFile("in.txt")

	// Preprocess
	input_split := strings.Split(string(input), "\n\n")
	rule_lines := strings.Split(string(input_split[0]), "\n")
	part_lines := strings.Split(string(input_split[1]), "\n")

	// Fill WorkFlows
	workFlows := map[string]*WorkFlow{}
	workFlows["A"] = &WorkFlow{"A", []Rule{}}
	workFlows["R"] = &WorkFlow{"R", []Rule{}}
	for _, line := range rule_lines {
		split := strings.SplitN(line, "{", 2)
		workFlows[split[0]] = &WorkFlow{split[0], make([]Rule, 0)}
	}
	// Add Rules
	for _, line := range rule_lines {
		parseRule(line, workFlows)
	}

	// Part 1
	sumP1 := 0
	for _, line := range part_lines {
		var x, m, a, s int
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d", &x, &m, &a, &s)
		sumP1 += makePart(Part{"x":x, "m":m, "a":a, "s":s}, *workFlows["in"])
	}
	fmt.Println("Sol 1:", sumP1)


	// Part 2
	sumP2 := makePartRanges(PartRanges{
		"x": {1, 4000}, 
		"m": {1, 4000}, 
		"a": {1, 4000}, 
		"s": {1, 4000}}, 
		*workFlows["in"],
	)
	fmt.Println("Sol 2:", sumP2)
}
