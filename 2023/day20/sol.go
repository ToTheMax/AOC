package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Broadcaster int = 0
	FlipFlop    int = 1
	Conjunction int = 2
	DeadEnd		int = 3
)


type Module struct {
	name string
	moduleType int
	destinations []*Module
	on bool
	memory map[string]bool
	destinations_str []string
}

type State struct {
	fromModule *Module
	toModule *Module
	highPulse bool
}

// Broadcast sends the given pulse to all of its destination modules.
//
// Flipflop initial -> off
// Flipflow high -> ignore
// Flipflop low -> flips between on and off and sends a high or low pulse 
//
// Inv (Conjunctio) remembers the type of the most recent pulse received from each of their connected input modules
// Inv intial -> low pulse
// Inv high/low -> update memory, if all high -> low pulse, else high pulse
//
//
// button sends a single low pulse is sent directly to the broadcaster module.
// Pulses are always processed in the order they are sent (BFS)


func pressButton(modules map[string]*Module, highPulseCount *int, lowPulseCount *int) {
	queue := []*State{}
	
	queue = append(queue, &State{&Module{"button", Broadcaster, make([]*Module, 0), false, map[string]bool{}, []string{}}, modules["broadcaster"], false})
	for len(queue) > 0 {
		current := queue[0]
		// fmt.Println(current.fromModule.name, current.highPulse, current.toModule.name)
		queue = queue[1:]

		if current.highPulse {
			*highPulseCount++
		} else {
			*lowPulseCount++
		}


		switch current.toModule.moduleType {
		case Broadcaster:
			for _, dest := range current.toModule.destinations {
				queue = append(queue, &State{current.toModule, dest, current.highPulse})
			}
			
		case FlipFlop:
			if !current.highPulse {
				current.toModule.on = !current.toModule.on
				for _, dest := range current.toModule.destinations {
					queue = append(queue, &State{current.toModule, dest, current.toModule.on})
				}
			}
		case Conjunction:
			current.toModule.memory[current.fromModule.name] = current.highPulse
			allHigh := true
			// fmt.Println(current.toModule.memory)
			for _, val := range current.toModule.memory {
				allHigh = allHigh && val
			}
			// fmt.Println("allHigh", allHigh)
			for _, dest := range current.toModule.destinations {
				queue = append(queue, &State{current.toModule, dest, !allHigh})
			}
		case DeadEnd:
			// Do nothing
		}
	}
}


func main() {
	input, _ := os.ReadFile("in.txt")
	lines := strings.Split(string(input), "\n")

	modules := make(map[string]*Module)

	// Fill modules
	for _, line := range lines{
		split := strings.SplitN(line, " -> ", 2)
		name := split[0]
		dests := strings.Split(split[1], ", ")
		
		switch name[0] {
		case '%':
			modules[name[1:]] = &Module{name, FlipFlop, make([]*Module, 0), false, map[string]bool{}, dests}
		case '&':
			modules[name[1:]] = &Module{name, Conjunction, make([]*Module, 0), false,  map[string]bool{}, dests}
		case 'b':
			modules["broadcaster"] = &Module{name, Broadcaster, make([]*Module, 0), false,  map[string]bool{}, dests}
		}
	}

	// Fill destinations and memory
	for _, module := range modules{
		for _, dest := range module.destinations_str {
			if modules[dest] == nil {
				modules[dest] = &Module{dest, DeadEnd, make([]*Module, 0), false, map[string]bool{}, []string{}}
			}
			module.destinations = append(module.destinations, modules[dest])
			if modules[dest].moduleType == Conjunction {
				modules[dest].memory[module.name] = false
			}
		}
	}

	// Press button
	highPulseCount, lowPulseCount  := 0, 0
	for i:=0; i<1000; i++ {
		pressButton(modules, &highPulseCount, &lowPulseCount)
	}
	fmt.Println("Sol 1:", highPulseCount* lowPulseCount)
}
