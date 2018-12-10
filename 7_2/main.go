package main

import (
	"bufio"
	"bytes"
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

	steps := Steps{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var req, target string
		_, err := fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &req, &target)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}

		// creating both steps if not created
		if _, ok := steps[target[0]]; !ok {
			steps[target[0]] = &Step{Reqs: []byte{}}
		}
		if _, ok := steps[req[0]]; !ok {
			steps[req[0]] = &Step{Reqs: []byte{}}
		}

		// add dependency
		steps[target[0]].Reqs = append(steps[target[0]].Reqs, req[0])
	}

	index := 1
	duraitons := map[byte]int{}
	for i := byte('A'); i <= byte('Z'); i++ {
		// duraitons[i] = index
		duraitons[i] = 60 + index
		index++
	}

	releaseSchedule := map[byte]int{}
	// workerAvailableAt := WorkersSchedule{1: 0, 2: 0}

	workerAvailableAt := WorkersSchedule{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	response := []byte{}
	for {
		fmt.Println("-------------------")
		fmt.Printf("Release schedule: %v\n", releaseSchedule)
		// steps.print()

		depsToClear := workerAvailableAt.doReleaseSchedule(&releaseSchedule)
		for _, id := range depsToClear {
			fmt.Printf("clearDependencies for %c\n", id)
			steps.clearDependencies(id)
		}

		nextStep, err := steps.findNextStep()
		if err != nil {
			fmt.Printf("next step not found, result so far: %s\n", string(response))
			err = workerAvailableAt.DoIdle()
			if err != nil {
				fmt.Printf("No idling result so far: %s\n", string(response))
				return
			}

			continue
		}
		response = append(response, nextStep)
		delete(steps, nextStep)

		releaseTime := workerAvailableAt.work(0, duraitons[nextStep])
		releaseSchedule[nextStep] = releaseTime
		workerAvailableAt.print()
	}

	// steps.print()
	// steps.clearDependencies(byte('B'))
	// steps.print()
}

type Step struct {
	Reqs []byte
}

type Steps map[byte]*Step

func (steps *Steps) print() {
	for id, step := range *steps {
		fmt.Printf("instruction: %c reqs: %s\n", id, string(step.Reqs))
	}
}

func (steps *Steps) findNextStep() (nextStep byte, err error) {
	nextStep = 0
	fmt.Printf("Possible steps:")
	for id, step := range *steps {
		if len(step.Reqs) == 0 {
			fmt.Printf("%c", id)
			if nextStep == 0 || id < nextStep {
				nextStep = id
			}
		}
	}

	fmt.Printf("; chosen: %c\n", nextStep)
	if nextStep == 0 {
		return 0, fmt.Errorf("next step not found")
	}

	return
}

func (steps *Steps) clearDependencies(stepToExecute byte) {
	for id, _ := range *steps {
		if pos := bytes.IndexByte((*steps)[id].Reqs, stepToExecute); pos != -1 {
			(*steps)[id].Reqs = append((*steps)[id].Reqs[:pos], (*steps)[id].Reqs[pos+1:]...)
		}
	}
}

type WorkersSchedule map[int]int

func (w *WorkersSchedule) findFirstAvailable() (earliestId int) {
	earliestId = -1
	earliestTime := 0
	for id, time := range *w {
		if earliestId == -1 || time < earliestTime {
			earliestId = id
			earliestTime = time
		}
	}
	return earliestId
}

func (w *WorkersSchedule) curTime() (earliestTime int) {
	earliestTime = -1
	for _, time := range *w {
		if earliestTime == -1 || time <= earliestTime {
			earliestTime = time
		}
	}
	return earliestTime
}

func (w *WorkersSchedule) work(from, duration int) (endTime int) {
	id := w.findFirstAvailable()
	if (*w)[id] > from {
		from = (*w)[id]
	}
	(*w)[id] = from + duration
	return (*w)[id]
}

func (w *WorkersSchedule) print() {
	for id, time := range *w {
		fmt.Printf("\t%d:%d", id, time)
	}
	fmt.Println("")
}

func (w *WorkersSchedule) doReleaseSchedule(releaseSchedule *map[byte]int) (ids []byte) {
	curTime := w.curTime()
	fmt.Printf("doReleaseSchedule curTime: %d\n", curTime)
	for id, time := range *releaseSchedule {
		if time <= curTime {
			ids = append(ids, id)
			delete(*releaseSchedule, id)
		}
	}
	return
}

// here we'll push timeline to the next
func (w *WorkersSchedule) DoIdle() (err error) {
	fmt.Printf("DoIdle start: %v\n", *w)
	minTime := 4000000000
	secondMinTime := 4000000000

	for _, time := range *w {
		if time < minTime {
			minTime = time
		}
	}

	for _, time := range *w {
		if time > minTime && time < secondMinTime {
			secondMinTime = time
		}
	}

	if secondMinTime == 4000000000 {
		return fmt.Errorf("")
	}

	for id, time := range *w {
		if time == minTime {
			(*w)[id] = secondMinTime
		}
	}

	fmt.Printf("DoIdle done: %v\n", *w)
	return
}
