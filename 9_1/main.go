package main

import (
	"fmt"
)

func main() {
	// task
	maxMarbleValue := 7143100
	playersNum := 476

	// first test
	// maxMarbleValue := 25
	// playersNum := 9

	// test 1
	// maxMarbleValue := 1618
	// playersNum := 10

	players := players{
		count:         playersNum,
		currentPlayer: 1,
		scores:        make(map[int]int),
	}

	for i := 1; i <= players.count; i++ {
		players.scores[i] = 0
	}

	// initialize the first marble
	m := marbleInCircle{value: 0}
	m.cwMarble = &m
	m.ccwMarble = &m

	sumScores := 0

	curMarble := &m
	curMarbleValue := 1
	for curMarbleValue <= maxMarbleValue {
		// fmt.Printf("adding the marble #%d\n", curMarbleValue)
		if curMarbleValue%23 > 0 {
			curMarble = curMarble.stepCw(1)
			curMarble = curMarble.addMarbleCw(curMarbleValue)
		} else {
			var addScore int
			curMarble = curMarble.stepCcw(7)
			curMarble, addScore = curMarble.removeCurrentMarbleAndCw()
			players.ScoreCurrentPlayer(curMarbleValue + addScore)
			sumScores += curMarbleValue + addScore
			// fmt.Printf("scoring the marble #%d\n", curMarbleValue)
		}

		players.Next()
		curMarbleValue++
	}

	// m.printState()
	// fmt.Printf("scores: %v\n", players)
	// fmt.Printf("currentMarble: %v\n", curMarble.value)
	fmt.Printf("max score: %d, sumScores: %d/%d\n", players.maxScore(), players.sumScores(), sumScores)
}

type marbleInCircle struct {
	value     int
	cwMarble  *marbleInCircle
	ccwMarble *marbleInCircle
}

func (m *marbleInCircle) stepCw(steps int) (newCurMarble *marbleInCircle) {
	newCurMarble = m
	for i := 0; i < steps; i++ {
		newCurMarble = newCurMarble.cwMarble
	}
	return
}

func (m *marbleInCircle) stepCcw(steps int) (newCurMarble *marbleInCircle) {
	newCurMarble = m
	for i := 0; i < steps; i++ {
		newCurMarble = newCurMarble.ccwMarble
	}
	return
}

func (m *marbleInCircle) addMarbleCw(value int) (newCurMarble *marbleInCircle) {
	// inserting new marble
	newMarble := &marbleInCircle{
		value:     value,
		cwMarble:  m.cwMarble,
		ccwMarble: m,
	}
	// updating existing marbles
	m.cwMarble.ccwMarble = newMarble
	m.cwMarble = newMarble
	return newMarble
}

func (m *marbleInCircle) addMarbleCcw(value int) (newCurMarble *marbleInCircle) {
	// inserting new marble
	newMarble := &marbleInCircle{
		value:     value,
		cwMarble:  m,
		ccwMarble: m.ccwMarble,
	}
	// updating existing marbles
	m.ccwMarble.cwMarble = newMarble
	m.ccwMarble = newMarble
	return m
}

func (m *marbleInCircle) removeCurrentMarbleAndCw() (newCurMarble *marbleInCircle, value int) {
	m.ccwMarble.cwMarble = m.cwMarble
	m.cwMarble.ccwMarble = m.ccwMarble
	return m.cwMarble, m.value
}

func (m *marbleInCircle) printState() {
	fmt.Print("marbles: ")
	// storing current value
	c := m
	for {
		fmt.Printf("%d ", c.value)
		c = c.stepCw(1)
		if c.value == m.value {
			fmt.Println("")
			return
		}
	}
}

type players struct {
	count         int
	currentPlayer int
	scores        map[int]int
}

func (p *players) Next() int {
	p.currentPlayer++
	if p.currentPlayer > p.count {
		p.currentPlayer = 1
	}
	return p.currentPlayer
}

func (p *players) ScoreCurrentPlayer(value int) {
	p.scores[p.currentPlayer] += value
}

func (p *players) sumScores() (res int) {
	for _, score := range p.scores {
		res += score
	}
	return
}

func (p *players) maxScore() (max int) {
	for _, score := range p.scores {
		if score > max {
			max = score
		}
	}
	return
}

// type encryptedSequenceType []int

// var maxCalls = 10000

// func (sequence *encryptedSequenceType) parseNodes(numNodes, start, level int) (metadataSum, ends int) {
// 	fmt.Printf("%sparsing %d nodes starting from %d\n", strings.Repeat("\t", level), numNodes, start)
// 	maxCalls--
// 	if maxCalls < 0 {
// 		log.Fatalf("too deep\n")
// 	}
// 	cursor := start

// 	for i := 1; i <= numNodes; i++ {
// 		childMetadataSum := 0
// 		childNodes := (*sequence)[cursor]
// 		curNodeMetadataCount := (*sequence)[cursor+1]
// 		fmt.Printf("%s. parsing #%d node starting from %d: %d childs, %d metadata\n", strings.Repeat("\t", level), i, cursor, childNodes, curNodeMetadataCount)
// 		cursor += 2

// 		if childNodes != 0 {
// 			childMetadataSum, cursor = sequence.parseNodes(childNodes, cursor, level+1)
// 			fmt.Printf("%s+ returned next node position %d\n", strings.Repeat("\t", level), cursor)
// 		}
// 		fmt.Printf("%s. parsing #%d node: childEnds %d\n", strings.Repeat("\t", level), i, cursor)

// 		metadataSum += childMetadataSum
// 		fmt.Printf("%s. metadata:", strings.Repeat("\t", level))
// 		for j := 0; j < curNodeMetadataCount; j++ {
// 			metadataSum += (*sequence)[cursor]
// 			fmt.Printf("+%d", (*sequence)[cursor])
// 			cursor++
// 		}
// 		fmt.Println("")
// 	}
// 	return metadataSum, cursor
// }
