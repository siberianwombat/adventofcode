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
	f, err := os.Open("input.txt")
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

	msum, ends := sample.parseNodes(1, 0)
	fmt.Printf("%d, %d\n", msum, ends)
}

type encryptedSequenceType []int

var maxCalls = 10

func (sequence *encryptedSequenceType) parseNodes(numNodes, start int) (metadataSum, ends int) {
	fmt.Printf("parsing %d nodes starting from %d\n", numNodes, start)
	maxCalls--
	if maxCalls < 0 {
		log.Fatalf("too deep\n")
	}
	cursor := start
	childEnds := 0

	for i := 1; i <= numNodes; i++ {
		childMetadataSum := 0
		childNodes := (*sequence)[cursor]
		curNodeMetadataCount := (*sequence)[cursor+1]
		cursor += 2
		fmt.Printf(". parsing #%d node: %d childs, %d metadata\n", i, childNodes, curNodeMetadataCount)
		if childNodes != 0 {
			childMetadataSum, childEnds = sequence.parseNodes(childNodes, cursor)
			fmt.Printf("+ returned next node position %d\n", childEnds)
			cursor = childEnds + 1
		}
		fmt.Printf(". parsing #%d node: childEnds %d\n", i, childEnds)

		metadataSum += childMetadataSum
		fmt.Printf(". metadata:")
		for j := 0; j < curNodeMetadataCount; j++ {
			metadataSum += (*sequence)[cursor+j]
			fmt.Printf("+%d", (*sequence)[cursor+j])
			ends = cursor + 1 + j
		}
		fmt.Println("")
		cursor = ends + 1
	}
	return metadataSum, cursor
}
