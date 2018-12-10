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
	guards := Guards{}
	var curId int
	var curFallMinute int
	maxMin := 0
	maxMinCount := 0
	maxMinsId := 0
	for s.Scan() {
		var y, m, d, h, min int
		var message1, message2 string
		_, err := fmt.Sscanf(s.Text(), "[%d-%d-%d %d:%d] %s %s", &y, &m, &d, &h, &min, &message1, &message2)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}

		switch message1 {
		case "Guard":
			_, err := fmt.Sscanf(message2, "#%d", &curId)
			if err != nil {
				log.Fatalf("failed reading: %s", err.Error())
			}
		case "falls":
			curFallMinute = min
		case "wakes":
			if _, ok := guards[curId]; !ok {
				guards[curId] = &GuardHistory{
					MaxMinCount: 0,
					// History:     []GuardHistoryItem{},
					Minutes: [60]int{},
				}
			}
			for i := curFallMinute; i < min; i++ {
				guards[curId].Minutes[i]++
				if guards[curId].Minutes[i] > guards[curId].MaxMinCount {
					guards[curId].MaxMinCount = guards[curId].Minutes[i]
					if guards[curId].MaxMinCount > maxMinCount {
						maxMinsId = curId
						maxMinCount = guards[curId].MaxMinCount
						maxMin = i
					}
				}
			}
		}
	}

	fmt.Printf("Max mins [%d] @%d Guard #%d\n", maxMinCount, maxMin, maxMinsId)

}

type GuardHistory struct {
	MaxMin      int
	MaxMinCount int
	Minutes     [60]int
}

type Guards map[int]*GuardHistory
