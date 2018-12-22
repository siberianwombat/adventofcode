package main

import "fmt"

func main() {
	serialNumber := 5177

	batteryGrid := grid{}

	// first cycle
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			batteryGrid[coords{x: x, y: y}] = battery{powerLevel: powerLevel(x, y, serialNumber)}
		}
	}

	var maxX, maxY, maxPowerLevel int
	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			curPowerLevel := batteryGrid.g3x3powerLevel(x, y)
			if curPowerLevel > maxPowerLevel {
				maxX = x
				maxY = y
				maxPowerLevel = curPowerLevel
			}
		}
	}

	fmt.Printf("%d:%d -> %d\n", maxX, maxY, maxPowerLevel)
	// fmt.Printf("sample result: %d\n", batteryGrid.g3x3powerLevel(21, 61))
	// fmt.Printf("sample result: %d\n", powerLevel(3, 5, serialNumber))
	// fmt.Printf("sample result: %d\n", powerLevel(122, 79, 57))
	// fmt.Printf("sample result: %d\n", powerLevel(217, 196, 39))
	// fmt.Printf("sample result: %d\n", powerLevel(101, 153, 71))
}

func powerLevel(x, y, serialNumber int) (powerLevel int) {
	rackId := x + 10
	powerLevel = rackId * y
	powerLevel += serialNumber
	powerLevel *= rackId
	powerLevel = (powerLevel / 100) % 10
	powerLevel = powerLevel - 5
	return
}

type coords struct {
	x int
	y int
}

type battery struct {
	powerLevel     int
	g3x3powerLevel int
}

type grid map[coords]battery

func (g *grid) g3x3powerLevel(x, y int) (powerLevel int) {
	if x < 1 || y < 1 || x > 300-2 || y > 300-2 {
		return 0
	}

	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			powerLevel += ((*g)[coords{x: x + dx, y: y + dy}]).powerLevel
		}
	}

	return
}
