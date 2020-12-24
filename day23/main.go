package main

import (
	"container/ring"
	"fmt"
)

func main() {
	part1()
	part2()
}

func select3(cupValue int, cups []int) ([]int, []int) {
	index := findIndex(cupValue, cups)
	selection := []int{}
	removeIndexes := []int{}
	for i := 1; i <= 3; i++ {
		removeIndex := (index + i) % len(cups)
		selection = append(selection, cups[removeIndex])
		removeIndexes = append(removeIndexes, removeIndex)
	}
	result := []int{}
	for idx, val := range cups {
		allDiff := 0
		for _, removeIdx := range removeIndexes {
			if idx != removeIdx {
				allDiff++
			}
		}
		if allDiff == len(selection) {
			result = append(result, val)
		}
	}
	return selection, result
}

func max(cups []int) int {
	max := 0
	for _, val := range cups {
		if val > max {
			max = val
		}
	}
	return max
}

func findIndex(value int, cups []int) int {
	for idx, current := range cups {
		if current == value {
			return idx
		}
	}
	return -1
}

func next(cupValue int, cups []int) int {
	minus := 1
	for {
		next := cupValue - minus
		if next < 0 {
			max := max(cups)
			return findIndex(max, cups)
		}
		index := findIndex(next, cups)
		if index >= 0 {
			return index
		}
		minus++
	}
}

func addAfter(cupIndex int, toAdd []int, cups []int) []int {
	result := make([]int, len(cups[:cupIndex+1]))
	copy(result, cups[:cupIndex+1])
	result = append(result, toAdd...)
	if cupIndex+len(toAdd)+1 < len(cups)+len(toAdd) {
		result = append(result, cups[cupIndex+1:]...)
	}
	return result
}

func part1() {
	cups := []int{4, 6, 7, 5, 2, 8, 1, 9, 3}
	moves := 100
	var threeAfter []int
	currentCup := cups[0]
	for i := 0; i < moves; i++ {
		threeAfter, cups = select3(currentCup, cups)
		nextIndex := next(currentCup, cups)
		cups = addAfter(nextIndex, threeAfter, cups)
		currentCup = cups[(findIndex(currentCup, cups)+1)%len(cups)]
	}
	fmt.Println(cups)
}

func nextCup(currentCup *ring.Ring, trioHead *ring.Ring, cupMap map[int]*ring.Ring) *ring.Ring {
	hasMap := make(map[int]bool)
	hasMap[trioHead.Value.(int)] = true
	hasMap[trioHead.Next().Value.(int)] = true
	hasMap[trioHead.Next().Next().Value.(int)] = true
	getMax := currentCup.Value.(int) == 1 ||
		(currentCup.Value.(int) == 2 && hasMap[1]) ||
		(currentCup.Value.(int) == 3 && hasMap[1] && hasMap[2]) ||
		(currentCup.Value.(int) == 4 && hasMap[1] && hasMap[2] && hasMap[3])
	if getMax {
		if !hasMap[1_000_000] {
			return cupMap[1_000_000]
		} else if !hasMap[999_999] {
			return cupMap[999_999]
		} else if !hasMap[999_998] {
			return cupMap[999_999]
		} else {
			return cupMap[999_997]
		}
	} else {
		next := currentCup.Value.(int) - 1
		for {
			if !hasMap[next] {
				return cupMap[next]
			}
			next--
		}
	}
}

func part2() {
	cups := []int{4, 6, 7, 5, 2, 8, 1, 9, 3}
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, i)
	}
	cupMap := make(map[int]*ring.Ring)
	cupRing := ring.New(len(cups))
	for _, c := range cups {
		cupMap[c] = cupRing
		cupRing.Value = c
		cupRing = cupRing.Next()
	}
	currentCup := cupMap[cups[0]]
	for i := 0; i < 10000000; i++ {
		trioHead := currentCup.Unlink(3)
		nextCup := nextCup(currentCup, trioHead, cupMap)
		nextCup.Link(trioHead)
		currentCup = currentCup.Next()
	}
	fmt.Println(cupMap[1].Next().Value, cupMap[1].Next().Next().Value,
		cupMap[1].Next().Value.(int)*cupMap[1].Next().Next().Value.(int))
}
