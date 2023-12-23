package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	field string
	leq   bool
	value int
	dest  *WorkFlow
}

type WorkFlow struct {
	name  string
	rules []Rule
	final bool
}

func main() {
	input, _ := os.ReadFile("in.txt")

	// Preprocess
	input_split := strings.Split(string(input), "\n\n")

	// Fill WorkFlows
	workFlows := map[string]*WorkFlow{}
	for _, line := range strings.Split(string(input_split[0]), "\n") {
		split1 := strings.SplitN(line, "{", 2)
		workFlows[split1[0]] = &WorkFlow{split1[0], make([]Rule, 0), false}
	}

	workFlows["A"] = &WorkFlow{"A", []Rule{}, true}
	workFlows["R"] = &WorkFlow{"A", []Rule{}, true}

	// Add Rules
	for _, line := range strings.Split(string(input_split[0]), "\n") {
		split1 := strings.SplitN(line, "{", 2)
		workFlowName := split1[0]
		rules := strings.Split(split1[1][:len(split1[1])-1], ",")

		for _, rule := range rules {
			split2 := strings.SplitN(rule, ":", 2)

			if len(split2) == 1 {
				fmt.Println(split2[0])
				workFlows[workFlowName].rules = append(
					workFlows[workFlowName].rules, Rule{"", false, 0, workFlows[split2[0]]},
				)
			} else {
				condition := split2[0]
				val, _ := strconv.Atoi(condition[3:])
				workFlows[workFlowName].rules = append(
					workFlows[workFlowName].rules, Rule{condition[:1], condition[1] == '<', val, workFlows[split2[1]]},
				)
			}
		}
		fmt.Println(workFlows)
	}

}
