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
	// polimer := []byte{'a', 'Z', 'z', 'A', 'Z', 'Z'}

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
			fmt.Printf("%s\n", string(newPolimer))
			fmt.Printf("len: %d\n", len(newPolimer))
			return
		} else {
			j++
			fmt.Printf("iteration: %d, len: %d\n", j, len(newPolimer))
			polimer = newPolimer
		}
	}

	// fmt.Printf("read %d bytes\n", len(polimer))
}

func doCollapse(a, b byte) bool {
	if (a < 91 && b < 91) || (a > 91 && b > 91) {
		// fmt.Printf(". same case %d %d\n", a, b)
		return false // same case
	}
	// fmt.Printf(". diff: %d-%d=%d, %d-%d=%d\n", a, b, a-b, b, a, b-a)

	return a-b == 32 || b-a == 32
}
