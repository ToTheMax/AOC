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
	final bool
}

type Part = map[string]int

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
	if workFlow.final {
		if workFlow.name == "A" {
			return part["x"] + part["m"] + part["a"] + part["s"]
		} else 	{
			return 0
		}
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

func main() {
	input, _ := os.ReadFile("in.txt")

	// Preprocess
	input_split := strings.Split(string(input), "\n\n")
	rule_lines := strings.Split(string(input_split[0]), "\n")
	part_lines := strings.Split(string(input_split[1]), "\n")

	// Fill WorkFlows
	workFlows := map[string]*WorkFlow{}
	for _, line := range rule_lines {
		split := strings.SplitN(line, "{", 2)
		workFlows[split[0]] = &WorkFlow{split[0], make([]Rule, 0), false}
	}

	workFlows["A"] = &WorkFlow{"A", []Rule{}, true}
	workFlows["R"] = &WorkFlow{"R", []Rule{}, true}

	// Add Rules
	for _, line := range rule_lines {
		parseRule(line, workFlows)
	}

	// 
	sumP1 := 0
	for _, line := range part_lines {
		var part = Part{}
		var x, m, a, s int
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d", &x, &m, &a, &s)
		part["x"] = x
		part["m"] = m
		part["a"] = a
		part["s"] = s
		sumP1 += makePart(part, *workFlows["in"])
	}
	fmt.Println("Sol 1:", sumP1)

}
