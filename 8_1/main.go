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

	msum, ends := sample.parseNodes(1, 0, 0)
	fmt.Printf("%d, %d\n", msum, ends)
}

type encryptedSequenceType []int

var maxCalls = 10000

func (sequence *encryptedSequenceType) parseNodes(numNodes, start, level int) (metadataSum, ends int) {
	fmt.Printf("%sparsing %d nodes starting from %d\n", strings.Repeat("\t", level), numNodes, start)
	maxCalls--
	if maxCalls < 0 {
		log.Fatalf("too deep\n")
	}
	cursor := start

	for i := 1; i <= numNodes; i++ {
		childMetadataSum := 0
		childNodes := (*sequence)[cursor]
		curNodeMetadataCount := (*sequence)[cursor+1]
		fmt.Printf("%s. parsing #%d node starting from %d: %d childs, %d metadata\n", strings.Repeat("\t", level), i, cursor, childNodes, curNodeMetadataCount)
		cursor += 2

		if childNodes != 0 {
			childMetadataSum, cursor = sequence.parseNodes(childNodes, cursor, level+1)
			fmt.Printf("%s+ returned next node position %d\n", strings.Repeat("\t", level), cursor)
		}
		fmt.Printf("%s. parsing #%d node: childEnds %d\n", strings.Repeat("\t", level), i, cursor)

		metadataSum += childMetadataSum
		fmt.Printf("%s. metadata:", strings.Repeat("\t", level))
		for j := 0; j < curNodeMetadataCount; j++ {
			metadataSum += (*sequence)[cursor]
			fmt.Printf("+%d", (*sequence)[cursor])
			cursor++
		}
		fmt.Println("")
	}
	return metadataSum, cursor
}
