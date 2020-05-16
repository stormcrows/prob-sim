package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const numPrisoners = 100
const numSelections = 50

func shuffle(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

func getRange(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

func findWithinSelection(x int, choices []int) bool {
	for i := 0; i < numSelections; i++ {
		if choices[i] == x {
			return true
		}
	}

	return false
}

func simulateRandomChoices(drawers []int) int {
	choices := make([]int, numPrisoners)
	copy(choices, drawers)

	for i := 0; i < numPrisoners; i++ {
		shuffle(choices)
		if !findWithinSelection(i, choices) {
			return 0
		}
	}

	return 1
}

func followNumbers(x int, drawers []int) bool {
	selection := drawers[x]
	for i := 0; i < numSelections; i++ {
		if selection == x {
			return true
		}
		selection = drawers[selection]
	}

	return false
}

func simulateFollowingNumbers(drawers []int) int {
	for i := 0; i < numPrisoners; i++ {
		if !followNumbers(i, drawers) {
			return 0
		}
	}

	return 1
}

func simulate(n int, f func(drawers []int) int) float64 {
	drawers := getRange(numPrisoners)
	wins := 0

	for i := 0; i < n; i++ {
		shuffle(drawers)
		wins += f(drawers)
	}
	prob := math.Round((float64(wins) / float64(n)) * 100.0)

	return prob
}

func main() {
	n := 10000
	rand.Seed(time.Now().UnixNano())

	pRandomChoices := simulate(n, simulateRandomChoices)
	pFollowinNumbers := simulate(n, simulateFollowingNumbers)

	msg := `
100 Prisoners problem simulation results
========================================

Number of simulations: %d

Scenario 1:
Prisoners make random choices:
	-> success rate: %.f%%

Scenario 2:
Prisoners open their own number first, then follow the boxes:
	-> success rate: %.f%%

It's much better to follow the numbers!

`

	fmt.Printf(msg, n, pRandomChoices, pFollowinNumbers)
}
