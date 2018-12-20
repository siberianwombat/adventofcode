package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input_task.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	sample := encryptedSequenceType{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, digit := range strings.Split(s.Text(), " ") {
			pint, err := strconv.ParseInt(digit, 10, 0)
			if err != nil {
				log.Fatalf("unparsed digit: %s", digit)
			}
			sample = append(sample, int(pint))
		}
	}

	msum, ends := sample.nodeValue(0, 0)
	fmt.Printf("%d, %d\n", msum, ends)
}

type encryptedSequenceType []int

var maxCalls = 10000

func (sequence *encryptedSequenceType) parseNodes(numNodes, start, level int) (ends int, childValues []int) {
	cursor := start
	childValues = []int{}
	var value int
	fmt.Printf("%sparseNodes:%d from %d\n", strings.Repeat("\t", level), numNodes, cursor)

	for i := 1; i <= numNodes; i++ {
		value, cursor = sequence.nodeValue(cursor, level)
		childValues = append(childValues, value)
	}
	return cursor, childValues
}

func (sequence *encryptedSequenceType) nodeValue(nodeStart int, level int) (value, nodeEnd int) {
	cursor := nodeStart
	childNodes := (*sequence)[cursor]
	nodeMetadataCount := (*sequence)[cursor+1]
	fmt.Printf("%sparse one Node from %d [childNodes:%d;metaCount:%d]\n", strings.Repeat("\t", level), cursor, childNodes, nodeMetadataCount)
	cursor += 2

	metadataSum := 0
	if childNodes == 0 {
		// summing meta values
		fmt.Printf("%ssumming meta: ", strings.Repeat("\t", level))
		for j := 0; j < nodeMetadataCount; j++ {
			metadataSum += (*sequence)[cursor]
			fmt.Printf("+%d", (*sequence)[cursor])
			cursor++
		}
		fmt.Println("")
	} else {
		// summing meta pointers
		var childValues []int
		cursor, childValues = sequence.parseNodes(childNodes, cursor, level+1)
		fmt.Printf("%s+ child node values: %v\n", strings.Repeat("\t", level), childValues)
		fmt.Printf("%s+ returned next node position %d\n", strings.Repeat("\t", level), cursor)
		fmt.Printf("%ssumming pointers: ", strings.Repeat("\t", level))
		for j := 0; j < nodeMetadataCount; j++ {
			pIndex := (*sequence)[cursor]
			cursor++
			if pIndex < 1 || pIndex > len(childValues) {
				fmt.Printf("x[%d]x", pIndex)
				continue
			}
			metadataSum += childValues[pIndex-1]
			fmt.Printf("+[%d]%d", pIndex, childValues[pIndex-1])
		}
		fmt.Println("")
	}

	fmt.Printf("%sreturning cursor: %d\n", strings.Repeat("\t", level), cursor)
	return metadataSum, cursor
}
