package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input_task.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var initialState string
	// rules := map[string]string{}

	pots := pots{
		potsState: map[int]bool{},
		rules: rules{
			0:  false,
			1:  false,
			2:  false,
			3:  false,
			4:  false,
			5:  false,
			6:  false,
			7:  false,
			8:  false,
			9:  false,
			10: false,
			11: false,
			12: false,
			13: false,
			14: false,
			15: false,
			16: false,
			17: false,
			18: false,
			19: false,
			20: false,
			21: false,
			22: false,
			23: false,
			24: false,
			25: false,
			26: false,
			27: false,
			28: false,
			29: false,
			30: false,
			31: false,
		},
	}

	for s.Scan() {
		if strings.Contains(s.Text(), "initial state") {
			_, err := fmt.Sscanf(s.Text(), "initial state: %s", &initialState)
			if err != nil {
				log.Fatalf("failed reading [%s]: %s", s.Text(), err.Error())
			}
			for i := 0; i < len(initialState); i++ {
				if initialState[i] == '#' {
					pots.potsState[i] = true
				} else {
					pots.potsState[i] = false
				}
			}
			pots.minMaxSet()
			fmt.Printf("state: %#v\n", pots)
		}

		if strings.Contains(s.Text(), " => ") {
			var pre, result string
			_, err := fmt.Sscanf(s.Text(), "%s => %s", &pre, &result)
			if err != nil {
				log.Fatalf("failed reading [%s]: %s", s.Text(), err.Error())
			}
			pots.rules[ruleToInt(pre)] = (result == "#")
		}

	}

	fmt.Printf("init state: %s, rules: %v \n", initialState, pots.rules)

	fmt.Printf("getRule [%d]: %d\n", 1, pots.getRule(1))

	pots.print()
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d:", i)
		pots.step()
		pots.print()
	}

	fmt.Printf("sum: %d\n", pots.indexSum())
}

type rules map[int]bool

type pots struct {
	minPot    int
	maxPot    int
	potsState map[int]bool
	rules     map[int]bool
}

func ruleToInt(ruleStr string) (rule int) {
	if len(ruleStr) != 5 {
		log.Fatalf("incorrect rule!")
	}

	binaryConv := map[int]int{
		0: 16,
		1: 8,
		2: 4,
		3: 2,
		4: 1,
	}

	for i := 0; i < 5; i++ {
		if ruleStr[i] == '#' {
			rule += binaryConv[i]
		}
	}

	return
}

func (p *pots) minMaxSet() {
	initialized := false
	for i, isSet := range p.potsState {
		if !initialized {
			p.minPot = i
			p.maxPot = i
			initialized = true
			continue
		}
		if isSet {
			if i < p.minPot {
				p.minPot = i
			}
			if i > p.maxPot {
				p.maxPot = i
			}
		}
	}
}

func (p *pots) getRule(position int) (rule int) {
	binaryConv := map[int]int{
		-2: 16,
		-1: 8,
		0:  4,
		1:  2,
		2:  1,
	}

	// if position < p.minPot || position > p.maxPot {
	// 	log.Fatalf("outside of range,  should me %d < %d < %d", p.minPot, position, p.maxPot)
	// }

	for i := -2; i <= 2; i++ {
		if ok, value := p.potsState[position+i]; ok && value {
			rule += binaryConv[i]
		}
	}
	return
}

func (p *pots) print() {
	fmt.Printf("%d:%d:%d\t", p.minPot, p.maxPot, p.indexSum())
	for i := p.minPot; i <= p.maxPot; i++ {
		if p.potsState[i] {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println("")
}

func (p *pots) checksum() (s []byte) {
	for i := p.minPot; i <= p.maxPot; i++ {
		if p.potsState[i] {
			s = append(s, '#')
		} else {
			s = append(s, '.')
		}
	}
	return
}

func (p *pots) step() {
	newState := map[int]bool{}
	for i := p.minPot - 5; i <= p.maxPot+5; i++ {
		ruleInt := p.getRule(i)
		if p.rules[ruleInt] {
			newState[i] = true
		}
	}
	p.potsState = newState
	p.minMaxSet()
}

func (p *pots) indexSum() (sum int) {
	for i := p.minPot; i <= p.maxPot; i++ {
		if p.potsState[i] {
			sum += i
		}
	}
	return
}
