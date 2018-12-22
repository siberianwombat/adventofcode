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

	var maxX, maxY, maxSize, maxPowerLevel int
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			curPowerLevel, curSize := batteryGrid.g3x3powerLevel(x, y)
			if curPowerLevel > maxPowerLevel {
				maxX = x
				maxY = y
				maxSize = curSize
				maxPowerLevel = curPowerLevel
			}
		}
		fmt.Printf(".")
	}

	fmt.Printf("%d,%d,%d -> %d\n", maxX, maxY, maxSize, maxPowerLevel)
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

func (g *grid) g3x3powerLevel(x, y int) (maxPowerLevel, maxSize int) {
	if x < 1 || y < 1 || x > 300 || y > 300 {
		return
	}

	var powerLevel int
	for size := 1; x+size <= 301 && y+size <= 301; size++ {
		for dx := 0; dx < size-1; dx++ {
			powerLevel += ((*g)[coords{x: x + size - 1, y: y + dx}]).powerLevel
			powerLevel += ((*g)[coords{x: x + dx, y: y + size - 1}]).powerLevel
		}
		powerLevel += ((*g)[coords{x: x + size - 1, y: y + size - 1}]).powerLevel
		if powerLevel > maxPowerLevel {
			maxPowerLevel = powerLevel
			maxSize = size
		}
	}

	return
}
