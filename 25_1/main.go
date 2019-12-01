package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
}

type Star struct {
	x                 int
	y                 int
	z                 int
	z1                int
	name              string
	constellationName string // named after first star in constellation
}

type Stars []Star

func readFile(filename string) (stars Stars, cCount int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %s: %s", filename, err.Error())
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// stars = stars{}

	for s.Scan() {
		var x, y, z, z1 int
		_, err := fmt.Sscanf(s.Text(), "%d,%d,%d,%d", &x, &y, &z, &z1)
		if err != nil {
			log.Fatalf("failed reading: %s", err.Error())
		}

		s := Star{
			x:  x,
			y:  y,
			z:  z,
			z1: z1,
		}

		s.setName()

		oldConstellations := stars.getAllConstellations(s)

		if len(oldConstellations) > 1 {
			s.constellationName = stars.mergeConstellations(oldConstellations)
		} else {
			if len(oldConstellations) == 1 {
				s.constellationName = oldConstellations[0]
			} else {
				s.constellationName = s.name
			}
		}

		stars = append(stars, s)
	}

	cCount = stars.countConstellations()

	return
}

func sq(value int) (sqv float64) {
	vf := float64(value)
	return math.Pow(vf, 2)
}

func (c *Star) setName() {
	c.name = fmt.Sprintf("%d:%d:%d:%d", c.x, c.y, c.z, c.z1)
}

func (c Stars) getAllConstellations(newStar Star) (names []string) {
	constellationsHash := map[string]bool{}
	for _, star := range c {
		if starsDistance(star, newStar) <= 3 {
			constellationsHash[star.constellationName] = true
		}
	}
	if len(constellationsHash) == 0 {
		return
	}
	for name, _ := range constellationsHash {
		names = append(names, name)
	}
	return
}

func (c Stars) countConstellations() (count int) {
	constellationsHash := map[string]bool{}
	for _, star := range c {
		constellationsHash[star.constellationName] = true
	}
	fmt.Printf("All constellations: %+v\n", constellationsHash)
	return len(constellationsHash)
}

func (c Stars) mergeConstellations(names []string) (chosenname string) {
	for _, name := range names {
		if chosenname == "" {
			chosenname = name
			continue
		}
		for i, star := range c {
			if star.constellationName == name {
				c[i].constellationName = chosenname
			}
		}
	}

	return
}

func starsDistance(s1, s2 Star) (dist float64) {
	//return math.Sqrt(sq(s1.x-s2.x) + sq(s1.y-s2.y) + sq(s1.z-s2.z) + sq(s1.z1-s2.z1))
	return math.Abs(float64(s1.x-s2.x)) +
		math.Abs(float64(s1.y-s2.y)) +
		math.Abs(float64(s1.z-s2.z)) +
		math.Abs(float64(s1.z1-s2.z1))
}
