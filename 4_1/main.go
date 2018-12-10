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
	maxMins := 0
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
					Sum:     0,
					History: []GuardHistoryItem{},
				}
			}
			guards[curId].Sum += min - curFallMinute
			guards[curId].History = append(guards[curId].History, GuardHistoryItem{
				Start: curFallMinute,
				End:   min,
			})
			if guards[curId].Sum > maxMins {
				maxMins = guards[curId].Sum
				maxMinsId = curId
			}
		}
	}

	fmt.Printf("Max mins [%d] Guard #%d\n", maxMins, maxMinsId)

	// for id, historyP := range guards {
	// 	fmt.Printf("Guard #%d %+v\n", id, *historyP)
	// }

	maxMinute := 0
	maxMinuteCount := 0
	minutes := [60]int{}
	fmt.Printf("minutes %+v\n", minutes)
	for _, historyItem := range guards[maxMinsId].History {
		for i := historyItem.Start; i < historyItem.End; i++ {
			minutes[i]++
			if minutes[i] > maxMinuteCount {
				maxMinute = i
				maxMinuteCount = minutes[i]
			}
		}
	}
	fmt.Printf("minutes %+v\n", minutes)
	fmt.Printf("maxMinute: %d, maxMinuteCount: %d, multiplier: %d\n", maxMinute, maxMinuteCount, maxMinute*maxMinsId)
}

type GuardHistoryItem struct {
	Start int
	End   int
}

type GuardHistory struct {
	Sum     int
	History []GuardHistoryItem
}

type Guards map[int]*GuardHistory
