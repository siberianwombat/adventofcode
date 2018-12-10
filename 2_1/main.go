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

	var twos = 0
	var threes = 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		chars := map[byte]int{}
		str := s.Text()
		for i := 0; i < len(s.Text()); i++ {
			// fmt.Printf("str: %d\n", str[i])
			if _, ok := chars[str[i]]; !ok {
				chars[str[i]] = 0
			}
			chars[str[i]]++
		}

		has_two := false
		has_three := false
		for _, num := range chars {
			if num == 2 {
				has_two = true
			}
			if num == 3 {
				has_three = true
			}
		}

		if has_two {
			twos++
		}

		if has_three {
			threes++
		}

	}

	fmt.Printf("checksum: %d\n", twos*threes)
}
