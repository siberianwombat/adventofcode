package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var MaxDistance int = 10000

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

	// lets determine starting point
	centerX, centerY := points.regionCenter()
	// centerX, centerY = points.searchForCenter(centerX, centerY, points.cellTotalDistance(centerX, centerY))
	fmt.Printf("Center [%d:%d]\n", centerX, centerY)
	r := 1
	sq := 0
	for {
		dsq := points.circleInRegion(centerX, centerY, r, MaxDistance)
		if dsq == 0 {
			break
		}
		// fmt.Printf("radius: %d, squares in: %d\n", r, dsq)
		sq += dsq
		r++
	}

	fmt.Printf("MAX sq: %d\n", sq+1) // count the  point space itself

	// fmt.Printf("sample point 4,3 distance: %d\n", points.cellTotalDistance(4, 3))
	// points.print(0, 0, 10, 10, MaxDistance)
}

type Coordinate struct {
	X int
	Y int
}

type Points map[int]*Coordinate

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

var directions = []Coordinate{
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
}

func (points *Points) searchForCenter(initX, initY, initDistance int) (centerX, centerY int) {
	minDistance := -1
	var minDistanceX, minDistanceY int
	for _, direction := range directions {
		distance := points.cellTotalDistance(initX+direction.X, initY+direction.Y)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			minDistanceX = initX + direction.X
			minDistanceY = initY + direction.Y
		}
	}

	if minDistance < initDistance {
		fmt.Printf("distance: %d\n", minDistance)
		return points.searchForCenter(minDistanceY, minDistanceX, minDistance)
	}

	return minDistanceX, minDistanceY
}

func (points *Points) regionCenter() (x, y int) {
	minX := -1
	minY := -1
	maxX := -1
	maxY := -1
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

	x = (maxX - minX) / 2
	y = (maxY - minY) / 2

	return
}

func (points *Points) cellTotalDistance(x, y int) (dist int) {
	for _, coords := range *points {
		dist += distance(x, y, coords.X, coords.Y)
	}
	return
}

func (points *Points) circleInRegion(x, y, r, maxDistance int) (sq int) {
	circleCoordsMap := circle(x, y, r)
	for circleCoords, _ := range circleCoordsMap {
		if points.cellTotalDistance(circleCoords.X, circleCoords.Y) < maxDistance {
			sq++
		}
	}
	return
}

func (points *Points) print(minX, minY, maxX, maxY, maxDistance int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			char := "."
			if points.cellTotalDistance(x, y) < maxDistance {
				char = "#"
			}

			for _, coords := range *points {
				if x == coords.X && y == coords.Y {
					char = "X"
				}
			}

			fmt.Printf("%s", char)
		}
		fmt.Println("")
	}
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
