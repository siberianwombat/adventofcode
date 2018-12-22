package main

import (
	"fmt"
	"strconv"
)

func main() {
	afterRecipies := 846601
	kitchen := kitchen{
		recipies:          recipies{3, 7},
		firstElfPosition:  0,
		secondElfPosition: 1,
	}
	recipiesCount := 0

	kitchen.print()

	for {
		kitchen.MakeRecipe()
		recipiesCount++
		fmt.Printf("MakeRecipe #%d\n", recipiesCount)
		kitchen.firstElfPosition = kitchen.Move(kitchen.firstElfPosition)
		kitchen.secondElfPosition = kitchen.Move(kitchen.secondElfPosition)
		// kitchen.print()
		if found, after := kitchen.CheckRecipies(afterRecipies); found {
			fmt.Printf("found after %d recipies\n", after)
			return
		}

		// if recipiesCount > 20 {
		// 	return
		// }

	}

	// fmt.Printf("%v\n", kitchen.recipies[afterRecipies:afterRecipies+10])

	return
}

type recipies []int

type kitchen struct {
	recipies          recipies
	firstElfPosition  int
	secondElfPosition int
}

func (k *kitchen) print() {
	for i, rating := range k.recipies {
		decoratedRating := fmt.Sprintf("%d", rating)
		if i == k.firstElfPosition {
			decoratedRating = "(" + decoratedRating + ")"
		} else {
			if i == k.secondElfPosition {
				decoratedRating = "[" + decoratedRating + "]"
			} else {
				decoratedRating = " " + decoratedRating + " "
			}
		}
		fmt.Printf("%s ", decoratedRating)
	}
	fmt.Println()
}

func (k *kitchen) MakeRecipe() {
	newRecipeInt := k.recipies[k.firstElfPosition] + k.recipies[k.secondElfPosition]
	newRecipeString := strconv.Itoa(newRecipeInt)
	for i := 0; i < len(newRecipeString); i++ {
		newRating, _ := strconv.Atoi(string(newRecipeString[i]))
		k.recipies = append(k.recipies, newRating)
	}
}

func (k *kitchen) Move(curPos int) (newPos int) {
	movements := k.recipies[curPos]
	for i := 1; i <= movements+1; i++ {
		curPos++
		if curPos >= len(k.recipies) {
			curPos = 0
		}
	}
	return curPos
}

func (k *kitchen) CheckRecipies(Sequence int) (found bool, after int) {
	SequenceLen := 6

	found = false
	if len(k.recipies) < 7 {
		return
	}

	lastIndex := len(k.recipies) - 1

	// always check 2 last possibilities only
	firstPretendent := 0
	pow := 1
	for i := 0; i <= SequenceLen-1; i++ {
		firstPretendent += k.recipies[lastIndex-i] * (pow)
		pow *= 10

	}

	secondPretendent := 0
	pow = 1

	for i := 0; i <= SequenceLen-1; i++ {
		secondPretendent += k.recipies[lastIndex-i-1] * (pow)
		pow *= 10

	}

	fmt.Printf("checking %d and %d\n", firstPretendent, secondPretendent)

	if firstPretendent == Sequence {
		return true, len(k.recipies) - SequenceLen
	}

	if secondPretendent == Sequence {
		return true, len(k.recipies) - SequenceLen - 1
	}

	return
}
