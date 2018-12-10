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

	codes := map[int]string{}
	i := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		// chars := map[byte]int{}
		codes[i] = s.Text()
		i++

		// for i := 0; i < len(s.Text()); i++ {
		// 	// fmt.Printf("str: %d\n", str[i])
		// 	if _, ok := chars[str[i]]; !ok {
		// 		chars[str[i]] = 0
		// 	}
		// 	chars[str[i]]++
		// }

		// has_two := false
		// has_three := false
		// for _, num := range chars {
		// 	if num == 2 {
		// 		has_two = true
		// 	}
		// 	if num == 3 {
		// 		has_three = true
		// 	}
		// }

		// if has_two {
		// 	twos++
		// }

		// if has_three {
		// 	threes++
		// }

	}

	for i := 0; i < len(codes); i++ {
		for j := i + 1; j <= len(codes); j++ {
			if len(codes[i]) != len(codes[j]) {
				continue
			}

			smatched := []byte{}
			nmatches := 0
			for s := 0; s < len(codes[i]); s++ {
				if codes[i][s] != codes[j][s] {
					nmatches++
				} else {
					smatched = append(smatched, codes[i][s])
				}
			}

			if nmatches == 1 {
				fmt.Printf("Found! %s - %s, common:[%s]\n", codes[i], codes[j], string(smatched))
				return
			}

			fmt.Printf(". %s - %s\n", codes[i], codes[j])
		}
	}

	fmt.Printf("checksum: \n")
}
