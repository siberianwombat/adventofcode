package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	var polimer []byte
	s := bufio.NewScanner(f)
	for s.Scan() {
		polimer = s.Bytes()
	}

	// fmt.Printf("A Z %t\n", doCollapse('A', 'Z'))
	// fmt.Printf("a z %t\n", doCollapse('a', 'z'))
	// fmt.Printf("a Z %t\n", doCollapse('a', 'Z'))
	// fmt.Printf("a a %t\n", doCollapse('a', 'a'))
	// fmt.Printf("Z Z %t\n", doCollapse('Z', 'Z'))
	// fmt.Printf("z Z %t\n", doCollapse('z', 'Z'))
	// polimer = []byte("dabAcCaCBAcCcaDA")

	for i := byte('a'); i <= byte('z'); i++ {
		fmt.Printf("wo %d collapsed: %d\n", i, collapsePolimer(removeChar(polimer, i)))
	}

}

func removeChar(polimer []byte, char1 byte) (clearedPolimer []byte) {
	var char2 byte
	if char1 <= 'Z' {
		char2 = char1 + 32
	} else {
		char2 = char1 - 32
	}

	for i := 0; i < len(polimer); i++ {
		if polimer[i] != char1 && polimer[i] != char2 {
			clearedPolimer = append(clearedPolimer, polimer[i])
		}
	}

	return
}

func collapsePolimer(polimer []byte) (collapsedLength int) {
	var clean bool
	j := 0
	for {
		clean = true
		newPolimer := []byte{}
		for i := 1; i < len(polimer); i++ {
			if doCollapse(polimer[i], polimer[i-1]) {
				clean = false
				i++
			} else {
				newPolimer = append(newPolimer, polimer[i-1])
				if i == len(polimer)-1 {
					newPolimer = append(newPolimer, polimer[i])
				}
			}
		}
		if clean {
			// fmt.Printf("%s\n", string(newPolimer))
			// fmt.Printf("len: %d\n", len(newPolimer))
			return len(newPolimer)
		} else {
			j++
			// fmt.Printf("iteration: %d, len: %d\n", j, len(newPolimer))
			polimer = newPolimer
		}
	}
}

func doCollapse(a, b byte) bool {
	if (a < 91 && b < 91) || (a > 91 && b > 91) {
		return false // same case
	}
	return a-b == 32 || b-a == 32
}
