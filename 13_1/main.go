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

	tracksMap := tracks{}
	carts := carts{}
	y := 0
	for s.Scan() {
		mapLine := s.Text()
		// looking for carts
		for x := 0; x < len(mapLine); x++ {
			if mapLine[x] == '>' || mapLine[x] == '<' || mapLine[x] == '^' || mapLine[x] == 'v' {
				carts = append(carts, cart{
					x:         x,
					y:         y,
					facing:    mapLine[x],
					behaviour: "left",
				})
				_, _, trackChar := cartFacing(mapLine[x])
				mapLine = replaceN(mapLine, x, trackChar)
			}
		}
		tracksMap[y] = mapLine
		y++
	}

	tracksMap.print(&carts)

	maxCycles := 2000
	for {
		sort.Sort(carts)
		fmt.Printf("carts: %#v\n", carts)
		for i := range carts {
			carts[i].cartMove()
			terrain := tracksMap[carts[i].y][carts[i].x]
			carts[i].cartPostMove(terrain)
			collided := carts.checkCollision(i)
			if collided {
				fmt.Printf("Collided: %d:%d\n", carts[i].x, carts[i].y)
				return
			}
		}
		tracksMap.print(&carts)

		maxCycles--
		if maxCycles <= 0 {
			return
		}
	}
	// doing steps
}

type tracks map[int]string

type cart struct {
	x         int
	y         int
	facing    byte
	behaviour string
}

func (t *tracks) print(c *carts) {
	for y := 0; y < len(*t); y++ {
		mapLine := (*t)[y]
		for _, cart := range *c {
			if cart.y == y {
				mapLine = replaceN(mapLine, cart.x, cart.facing)
			}
		}
		fmt.Printf("%s\n", mapLine)
	}
}

func cartFacing(s byte) (dx, dy int, track byte) {
	switch s {
	case '<':
		return -1, 0, '-'
	case '>':
		return 1, 0, '-'
	case '^':
		return 0, -1, '|'
	case 'v':
		return 0, 1, '|'
	}
	log.Fatalf("not a cart symbol: %T", s)
	return
}

func (c *cart) cartMove() {
	dx, dy, _ := cartFacing(c.facing)
	c.x += dx
	c.y += dy
}

func (c *cart) cartPostMove(terrain byte) {
	switch terrain {
	case '+':
		switch c.behaviour {
		case "left":
			c.turnLeft()
		case "right":
			c.turnRight()
		}
		c.nextBehaviour()
	case '/':
		switch c.facing {
		case '<', '>':
			c.turnLeft()
		case '^', 'v':
			c.turnRight()
		}
	case '\\':
		switch c.facing {
		case '<', '>':
			c.turnRight()
		case '^', 'v':
			c.turnLeft()
		}
	default: // no specific behavoir

	}
}

func (c *cart) turnLeft() {
	switch c.facing {
	case '<':
		c.facing = 'v'
	case '>':
		c.facing = '^'
	case '^':
		c.facing = '<'
	case 'v':
		c.facing = '>'
	}
}

func (c *cart) turnRight() {
	switch c.facing {
	case '<':
		c.facing = '^'
	case '>':
		c.facing = 'v'
	case '^':
		c.facing = '>'
	case 'v':
		c.facing = '<'
	}
}

func (c *cart) nextBehaviour() {
	switch c.behaviour {
	case "left":
		c.behaviour = "straight"
	case "straight":
		c.behaviour = "right"
	case "right":
		c.behaviour = "left"
	default:
		log.Fatalf("unknown cart behaviour: [%s]", c.behaviour)
	}
}

type carts []cart

// sort pattern: to determine which cart goes first
func (s carts) Len() int      { return len(s) }
func (s carts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s carts) Less(i, j int) bool {
	if s[i].y < s[j].y {
		return true
	}
	if s[i].y > s[j].y {
		return false
	}
	return s[i].x < s[j].x
}

func sortByOrder(carts carts) (sortedCarts carts) {
	sort.Sort(carts)
	return carts
}

func (s carts) checkCollision(cartID int) (collided bool) {
	collided = false
	for i := range s {
		if i == cartID {
			continue
		}
		if s[i].x == s[cartID].x && s[i].y == s[cartID].y {
			return true
		}
	}
	return
}

func replaceN(s string, n int, c byte) (resString string) {
	if n < 0 || n >= len(s) {
		log.Fatalf("n=%d outside of range len(s)=%d", n, len(s))
	}
	resString = s[:n] + string(c) + s[n+1:]
	return
}
