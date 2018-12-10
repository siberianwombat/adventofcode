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

	// chart := make(map[int]map[int]int)
	chart := [1000][1000]int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			chart[i][j] = 0
		}
	}

	unclaimed := map[int]bool{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, left, top, width, height int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}
		fmt.Printf("%d %d %d %d %d\n", id, left, top, width, height)
		overlapped := false
		for i := left; i < left+width; i++ {
			for j := top; j < top+height; j++ {
				if chart[i][j] != 0 {
					delete(unclaimed, chart[i][j])
					overlapped = true
				}
				chart[i][j] = id
			}
		}
		if !overlapped {
			unclaimed[id] = true
		}
	}

	// count := 0
	// for i := 0; i < 1000; i++ {
	// 	for j := 0; j < 1000; j++ {
	// 		// fmt.Printf("%d", chart[i][j])
	// 		if chart[i][j] > 1 {
	// 			count++
	// 		}
	// 	}
	// 	// fmt.Print("\n")
	// }

	fmt.Printf("square: %+v \n", unclaimed)
}
