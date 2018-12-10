package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// aaaaa.cccc
// aAaaa.cccc
// aaaddecccc
// aadddeccCc
// ..dDdeeccc
// bb.deEeecc
// bBb.eeee..
// bbb.eeefff
// bbb.eeffff
// bbb.ffffFf

// In this example, the areas of coordinates A, B, C, and F are infinite - while not shown here,
// their areas extend forever outside the visible grid. However, the areas of coordinates D and E
// are finite: D is closest to 9 locations, and E is closest to 17 (both including the coordinate's
// location itself). Therefore, in this example, the size of the largest area is 17.

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	points := Points{}
	id := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y int
		_, err := fmt.Sscanf(s.Text(), "%d, %d", &x, &y)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}
		points[id] = &Coordinate{
			X: x,
			Y: y,
		}
		id++
	}

	maxSquare := 0
	maxSquareId := 0
	infPoints := points.infinitePoints()
	// fmt.Printf("INF POINTS: %v\n", infPoints)

	for pointID, coords := range points {
		if _, ok := infPoints[pointID]; !ok {
			// fmt.Printf("Checking pointID %d\n", pointID)
			r := 1
			sumSq := 0
			for {
				sq := points.circleOwnerSquare(coords.X, coords.Y, r, pointID)
				if sq == 0 {
					break
				}
				sumSq += sq
				r++
			}
			fmt.Printf("%d: %d\n", pointID, sumSq)
			if sumSq > maxSquare {
				maxSquare = sumSq
				maxSquareId = pointID
			}
		}
	}

	fmt.Printf("MAX %d: %d\n", maxSquareId, maxSquare+1) // count the  point space itself
}

type Coordinate struct {
	X int
	Y int
}

type Points map[int]*Coordinate

func (points *Points) owner(x, y int) (pointId int) {
	minDistance := -1
	for curPointId, coords := range *points {
		distance := distance(x, y, coords.X, coords.Y)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			pointId = curPointId
		} else {
			if distance == minDistance {
				pointId = 0 // same distance
			}
		}
	}
	return
}

func circle(x, y, r int) (circle map[Coordinate]bool) {
	circle = make(map[Coordinate]bool)
	for dx := x - r; dx <= x+r; dx++ {
		circle[Coordinate{X: dx, Y: y - r}] = true
		circle[Coordinate{X: dx, Y: y + r}] = true
	}
	for dy := y - r; dy <= y+r; dy++ {
		circle[Coordinate{X: x - r, Y: dy}] = true
		circle[Coordinate{X: x + r, Y: dy}] = true
	}
	return
}

func (points *Points) infinitePoints() (infPoints map[int]bool) {
	infPoints = map[int]bool{}
	br := points.boundRectangle()
	fmt.Printf("boundRectange items: %d\n", len(br))
	for coords, _ := range br {
		infPoints[points.owner(coords.X, coords.Y)] = true
	}
	return
}

func (points *Points) boundRectangle() (circle map[Coordinate]bool) {
	minX := -1
	minY := -1
	maxX := -1
	maxY := -1
	circle = map[Coordinate]bool{}
	for _, coords := range *points {
		if minX == -1 || coords.X < minX {
			minX = coords.X
		}
		if maxX == -1 || coords.X > maxX {
			maxX = coords.X
		}
		if minY == -1 || coords.Y < minY {
			minY = coords.Y
		}
		if maxY == -1 || coords.Y > maxY {
			maxY = coords.Y
		}
	}

	// fmt.Printf("boundRectangle: %d, %d, %d, %d\n", minX, maxX, minY, maxY)

	// fmt.Printf("for: from %d to %d\n", minX-maxX+minX, maxX+maxX-minX)
	for dx := minX - maxX + minX; dx <= maxX+maxX-minX; dx++ {
		// fmt.Printf("dx: %d\n", dx)
		circle[Coordinate{X: dx, Y: minY - maxY + minY}] = true
		circle[Coordinate{X: dx, Y: maxY + maxY - minY}] = true
	}
	for dy := minY - maxY + minY; dy <= maxY+maxY-minY; dy++ {
		// fmt.Printf("dy: %d\n", dy)
		circle[Coordinate{X: minX - maxX + minX, Y: dy}] = true
		circle[Coordinate{X: maxX + maxX - minX, Y: dy}] = true
	}
	return
}

func (points *Points) circleOwnerSquare(x, y, r, ownerId int) (sq int) {
	circleCoordsMap := circle(x, y, r)
	for circleCoords, _ := range circleCoordsMap {
		if points.owner(circleCoords.X, circleCoords.Y) == ownerId {
			sq++
		}
	}
	return
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
