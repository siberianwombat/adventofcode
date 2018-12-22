package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input_task.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	cave := caveMap{
		cells:    map[coords]*mapCell{},
		warriors: warriors{},
	}

	y := 0
	for s.Scan() {
		mapLine := s.Text()
		cave.maxX = len(mapLine) - 1
		for x := 0; x < len(mapLine); x++ {
			switch mapLine[x] {
			case '#':
				cave.cells[coords{x: x, y: y}] = &mapCell{
					floor:      false,
					occupiedby: nil,
				}
			case '.':
				cave.cells[coords{x: x, y: y}] = &mapCell{
					floor:      true,
					occupiedby: nil,
				}
			case 'E':
				newWarrior := warrior{
					hitpoints: 200,
					wtype:     wElf,
					x:         x,
					y:         y,
				}
				cave.warriors = append(cave.warriors, &newWarrior)
				cave.cells[coords{x: x, y: y}] = &mapCell{
					floor:      true,
					occupiedby: &newWarrior,
				}
			case 'G':
				newWarrior := warrior{
					hitpoints: 200,
					wtype:     wGoblin,
					x:         x,
					y:         y,
				}
				cave.warriors = append(cave.warriors, &newWarrior)
				cave.cells[coords{x: x, y: y}] = &mapCell{
					floor:      true,
					occupiedby: &newWarrior,
				}
			}
		}
		y++
	}
	cave.maxY = y - 1

	cave.print()
	// game rounds
	rounds := 1
	for {
		fmt.Printf("\n -= Round %d =-\n", rounds)

		// actual game round
		sort.Sort(cave.warriors)
		for warriorID, warrior := range cave.warriors {
			adjWarrior := cave.adjancedEnemy(warrior.x, warrior.y, wenemy(warrior.wtype))
			if adjWarrior != nil {
				// fmt.Printf("%d:%d:%d has enemy near %d:%d:%d\n", warrior.wtype, warrior.x, warrior.y, adjWarrior.wtype, adjWarrior.x, adjWarrior.y)
			} else {
				cave.moveWarrior(warriorID)
				adjWarrior = cave.adjancedEnemy(cave.warriors[warriorID].x, cave.warriors[warriorID].y, wenemy(warrior.wtype))
			}

			if adjWarrior != nil {
				adjWarrior.hitpoints -= 3
			}
		}

		// @todo cleanup all warriors with hp <= 0

		rounds++
		// if rounds >= 50 {
		// 	fmt.Printf("quit by timeout\n")
		// 	return
		// }

		win, winner, hitpoints := cave.checkWarriorSides()
		cave.print()
		if win {
			fmt.Printf("WIN! %s:%d\n", winner, hitpoints)
			return
		}
	}

}

const wElf = 1
const wGoblin = 2

type warrior struct {
	hitpoints int
	wtype     int
	x         int
	y         int
}

type coords struct {
	x int
	y int
}

type mapCell struct {
	floor      bool
	occupiedby *warrior
	distance   int
}

type warriors []*warrior
type caveMap struct {
	maxX     int
	maxY     int
	cells    map[coords]*mapCell
	warriors warriors
}

func (c *caveMap) adjancedCoords(x, y int) (cs []coords) {
	cs = []coords{}
	if x > 0 {
		cs = append(cs, coords{x: x - 1, y: y})
	}

	if y > 0 {
		cs = append(cs, coords{x: x, y: y - 1})
	}

	if x < c.maxX {
		cs = append(cs, coords{x: x + 1, y: y})
	}

	if y < c.maxY {
		cs = append(cs, coords{x: x, y: y + 1})
	}

	return
}

func wenemy(wtype int) (etype int) {
	if wtype == wElf {
		return wGoblin
	}
	return wElf
}

// seeking for wtype enemy adjanced with least hitpoints
func (c *caveMap) adjancedEnemy(x, y, wtype int) (w *warrior) {
	for _, coord := range c.adjancedCoords(x, y) {
		if c.cells[coord].occupiedby != nil && c.cells[coord].occupiedby.wtype == wtype {
			if w == nil {
				w = c.cells[coord].occupiedby
				continue
			}
			// at least two enemies
			if c.cells[coord].occupiedby.hitpoints < w.hitpoints {
				w = c.cells[coord].occupiedby
				continue
			}
			if c.cells[coord].occupiedby.hitpoints > w.hitpoints {
				continue
			}
			// equal hitpoints, taking first reading-wise
			if c.cells[coord].occupiedby.y < w.y {
				w = c.cells[coord].occupiedby
				continue
			}
			if c.cells[coord].occupiedby.y > w.y {
				continue
			}
			if c.cells[coord].occupiedby.x < w.x {
				w = c.cells[coord].occupiedby
			}
		}
	}
	return
}

func (c *caveMap) print() {
	// fmt.Printf("all warriors: %#v\n", c.warriors)
	for y := 0; y <= c.maxY; y++ {
		warriors := []string{}
		for x := 0; x <= c.maxX; x++ {
			mapCell := c.cells[coords{x: x, y: y}]
			if mapCell.occupiedby != nil {
				if mapCell.occupiedby.wtype == wElf {
					fmt.Printf("E")
					warriors = append(warriors, fmt.Sprintf("%s:%d", "E", mapCell.occupiedby.hitpoints))
				} else {
					fmt.Printf("G")
					warriors = append(warriors, fmt.Sprintf("%s:%d", "G", mapCell.occupiedby.hitpoints))
				}
			} else {
				if mapCell.floor {
					fmt.Printf(".")
				} else {
					fmt.Printf("#")
				}
			}
		}
		// here we'll be outputing health
		fmt.Printf(" %v\n", warriors)
	}
}

type distanceMap map[coords]int // -1 inpassable, 0+ distance to possible target

func (w warriors) Len() int      { return len(w) }
func (w warriors) Swap(i, j int) { w[i], w[j] = w[j], w[i] }
func (w warriors) Less(i, j int) bool {
	if w[i].y < w[j].y {
		return true
	}
	if w[i].y > w[j].y {
		return false
	}
	return w[i].x < w[j].x
}

func (dm *distanceMap) print(maxX, maxY int) {
	maxDistance := 4000000
	fmt.Printf("\n---------  DMAP -----------\n")
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxY; x++ {
			if (*dm)[coords{x: x, y: y}] == -1 {
				fmt.Printf(".")
			} else {
				if (*dm)[coords{x: x, y: y}] == maxDistance {
					fmt.Printf(" ")
				} else {
					fmt.Printf("%d", (*dm)[coords{x: x, y: y}])
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n--------- /DMAP -----------\n")
}

func (c *caveMap) moveWarrior(warriorID int) {
	// fmt.Printf("initial warrior coordinates %d:%d\n", c.warriors[warriorID].x, c.warriors[warriorID].y)
	// fmt.Printf("all warriors: %#v\n", c.warriors)
	maxDistance := 4000000
	// #1 cleaning the distance table
	dMap := distanceMap{}
	for x := 0; x <= c.maxX; x++ {
		for y := 0; y <= c.maxY; y++ {
			cell := c.cells[coords{x: x, y: y}]
			if cell.floor && cell.occupiedby == nil {
				dMap[coords{x: x, y: y}] = maxDistance
			} else {
				dMap[coords{x: x, y: y}] = -1
			}
		}
	}

	// for all enemies determine
	inRangeSpaces := 0
	enemyType := wenemy(c.warriors[warriorID].wtype)
	for _, warrior := range c.warriors {
		if warrior.wtype == enemyType {
			for _, coord := range c.adjancedCoords(warrior.x, warrior.y) {
				if dMap[coord] != -1 {
					dMap[coord] = 0
					inRangeSpaces++
				}
			}
		}
	}
	if inRangeSpaces == 0 {
		return
	}

	modified := true
	for modified {
		modified = false
		for y := 1; y <= c.maxY-1; y++ {
			for x := 1; x <= c.maxX-1; x++ {
				curCellDistance := dMap[coords{x: x, y: y}]
				if curCellDistance == -1 {
					continue
				}
				// fmt.Printf("checking %d:%d = %d\n", x, y, curCellDistance)
				for _, coord := range c.adjancedCoords(x, y) {
					if dMap[coord] == -1 {
						continue
					}
					if dMap[coord]+1 < curCellDistance {
						// fmt.Printf(" from %v = %d\n", coord, dMap[coord]+1)
						// fmt.Printf("%d < %d\n", dMap[coord]+1, curCellDistance)
						curCellDistance = dMap[coord] + 1
						dMap[coords{x: x, y: y}] = curCellDistance
						modified = true
					}
				}
			}
		}
	}
	// dMap.print(c.maxX, c.maxY)

	// finding the best way
	minDistance := maxDistance
	newX := c.warriors[warriorID].x + 1
	newY := c.warriors[warriorID].y + 1
	for _, coord := range c.adjancedCoords(c.warriors[warriorID].x, c.warriors[warriorID].y) {
		if dMap[coord] == -1 || dMap[coord] == maxDistance {
			continue
		}
		if dMap[coord] < minDistance {
			minDistance = dMap[coord]
			newX = coord.x
			newY = coord.y
			continue
		}
		if dMap[coord] > minDistance {
			continue
		}
		if coord.y < newY {
			newX = coord.x
			newY = coord.y
			continue
		}
		if coord.y > newY {
			continue
		}
		if coord.x < newX {
			newX = coord.x
			newY = coord.y
		}
	}

	oldWx := c.warriors[warriorID].x
	oldWy := c.warriors[warriorID].y
	if minDistance < maxDistance {
		// changing map assignment
		oldCell := c.cells[coords{x: oldWx, y: oldWy}]
		warriorP := oldCell.occupiedby
		// fmt.Printf("warriorP, %#v\n", *warriorP)
		if warriorP == nil {
			log.Fatalf("no warrior at coordinates %d:%d!!!", oldWx, oldWy)
		}
		oldCell.occupiedby = nil
		c.cells[coords{x: oldWx, y: oldWy}] = oldCell

		oldCell = c.cells[coords{x: newX, y: newY}]
		oldCell.occupiedby = warriorP
		c.cells[coords{x: newX, y: newY}] = oldCell

		// // changing warrior coordinates
		(*warriorP).x = newX
		(*warriorP).y = newY
		// fmt.Printf("warriorP, %#v\n", *warriorP)
		// fmt.Printf("all warriors %#v\n", c.warriors)
	}

}

func (c *caveMap) removeWarriorById(warriorID int) {
	// fmt.Printf("removeWarriorById:%d\n", warriorID)
	pWarrior := c.warriors[warriorID]
	// remove from list
	c.warriors = append(c.warriors[:warriorID], c.warriors[warriorID+1:]...)
	// remove reverence to warrior
	c.cells[coords{x: pWarrior.x, y: pWarrior.y}].occupiedby = nil
}

func (c *caveMap) checkWarriorSides() (win bool, winSide string, winHitPoints int) {
	var elvesCount, giblinsCount, elvesHitpoints, giblinsGitpoints int

	toRemove := []int{}
	for i, warrior := range c.warriors {
		if warrior.hitpoints <= 0 {
			toRemove = append(toRemove, i)
			continue
		}
		if warrior.wtype == wElf {
			elvesCount++
			elvesHitpoints += warrior.hitpoints
		} else {
			giblinsCount++
			giblinsGitpoints += warrior.hitpoints
		}
	}
	for _, warriorID := range toRemove {
		c.removeWarriorById(warriorID)
	}

	if elvesCount == 0 {
		return true, "Goblins", giblinsGitpoints
	}
	if giblinsCount == 0 {
		return true, "Elves", elvesHitpoints
	}

	// fmt.Printf("checkWarriorSides: G%d E%d", giblinsCount, elvesCount)

	return false, "", 0
}
