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

	s := bufio.NewScanner(f)
	var nums []int
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}
		nums = append(nums, n)
	}

	var seen = map[int]bool{0: true}
	var sum = 0
	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				return
			} else {
				fmt.Printf(" . %d = %d\n", n, sum)
			}
			seen[sum] = true
		}
	}

	// fmt.Printf("%+v", seen)
	// fmt.Println(sum)
}
