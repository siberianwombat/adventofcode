package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input_task.txt")
	if err != nil {
		log.Fatalf("can't open file input.txt: %s", err.Error())
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	stars := stars{}

	for s.Scan() {
		var x, y, dx, dy int
		_, err := fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}
		// fmt.Printf("%d %d %d %d \n", x, y, dx, dy)
		stars = append(stars, star{
			x:  x,
			y:  y,
			dx: dx,
			dy: dy,
		})
	}

	curSpread := 4000000000

	i := 0
	for {
		minX, maxX, minY, maxY := stars.getSpread(i)
		newSpread := maxX - minX + maxY - minY
		fmt.Printf("step %d, %d:%d %d:%d\n", i, maxX, minX, minY, maxY)
		fmt.Printf("step %d, spread: %d, incr = %t\n", i, newSpread, newSpread > curSpread)
		if newSpread > curSpread {
			stars.print(i - 1)
			return
		}
		curSpread = newSpread
		i++
		// if i > 10 {
		// 	return
		// }
	}

	// fmt.Printf("stars: %+v \n", stars)
}

type star struct {
	x  int
	y  int
	dx int
	dy int
}

type stars []star

func (s *stars) getSpread(step int) (minX, maxX, minY, maxY int) {
	initialized := false
	for _, star := range *s {
		if initialized {
			if star.x+star.dx*step < minX {
				// fmt.Printf("minX changed: %d -> %d\n", minX, star.x+star.dx*step)
				minX = star.x + star.dx*step
			}
			if star.x+star.dx*step > maxX {
				maxX = star.x + star.dx*step
			}
			if star.y+star.dy*step < minY {
				minY = star.y + star.dy*step
			}
			if star.y+star.dy*step > maxY {
				maxY = star.y + star.dy*step
			}
		} else {
			minX = star.x + star.dx*step
			maxX = star.x + star.dx*step
			minY = star.y + star.dy*step
			maxY = star.y + star.dy*step
			initialized = true
		}
	}
	return
}

func (s *stars) print(step int) {
	minX, maxX, minY, maxY := s.getSpread(step)
	fmt.Printf("printing step %d, %d:%d, %d:%d\n", step, minX, maxX, minY, maxY)

	codes := map[int]string{}
	for s := minY; s <= maxY; s++ {
		codes[s] = strings.Repeat(" ", maxX-minX+1)
	}

	for _, star := range *s {
		y := star.y + star.dy*step
		x := star.x + star.dx*step - minX
		// fmt.Printf(". printing star on %d:%d\n", x, y)
		curString := codes[y]
		// fmt.Printf("curString before: %s\n", curString)
		curString = setChar(curString, x)
		// fmt.Printf("curString after : %s\n", curString)
		codes[y] = curString
	}

	for s := minY; s <= maxY; s++ {
		fmt.Printf("%s\n", codes[s])
	}

}

func setChar(s string, pos int) (resString string) {
	if pos < 0 || pos >= len(s) {
		log.Fatalf("range outside %s[%d]", s, pos)
	}

	if pos == 0 {
		return "*" + s[1:]
	}

	if pos == len(s)-1 {
		return s[:len(s)-1] + "*"
	}

	return s[:pos] + "*" + s[pos+1:]
}
