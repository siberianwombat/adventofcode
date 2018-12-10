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
		instructions = append(instructions, instruction{
			Req:    req,
			Target: target,
		})

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

	response := []byte{}
	for {
		fmt.Println("-------------------")
		steps.print()
		nextStep, err := steps.findNextStep()
		if err != nil {
			fmt.Printf("next step not found, result so far: %s\n", string(response))
			return
		}
		response = append(response, nextStep)
		steps.clearDependencies(nextStep)
		delete(steps, nextStep)
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

type instruction struct {
	Req    string
	Target string
}

var instructions []instruction
