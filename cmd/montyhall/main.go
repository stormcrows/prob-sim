package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const goat = "goat"
const car = "car"

func next(n int, l int) int {
	return (n + 1) % l
}

func simulateRemainScenario(doors []string) int {
	l := len(doors)
	playerChoice := rand.Intn(l)

	goatReveal := next(playerChoice, l)
	if doors[goatReveal] == car {
		goatReveal = next(goatReveal, l)
	}

	if doors[playerChoice] == car {
		return 1
	}

	return 0
}

func simulateSwitchScenario(doors []string) int {
	l := len(doors)
	playerChoice := rand.Intn(l)

	goatReveal := next(playerChoice, l)
	if doors[goatReveal] == car {
		goatReveal = next(goatReveal, l)
	}

	playerChoice = next(playerChoice, l)
	if playerChoice == goatReveal {
		playerChoice = next(playerChoice, l)
	}

	if doors[playerChoice] == car {
		return 1
	}

	return 0
}

func simulate(n int, f func(doors []string) int) float64 {
	doors := []string{goat, goat, car}
	l := len(doors)
	wins := 0

	for i := 0; i < n; i++ {
		rand.Shuffle(l, func(i, j int) { doors[i], doors[j] = doors[j], doors[i] })
		wins += f(doors)
	}
	prob := math.Round((float64(wins) / float64(n)) * 100.0)

	return prob
}

func main() {
	n := 1000000
	rand.Seed(time.Now().UnixNano())

	probNoSwitch := simulate(n, simulateRemainScenario)
	probSwitch := simulate(n, simulateSwitchScenario)

	msg := `
Monty Hall simulation results
=============================

Number of simulations: %d

Scenario 1:
Player remains with his original guess:
	-> success rate: %.f%%

Scenario 2:
Player decides to switch doors after goat reveal:
	-> success rate: %.f%%

It's much better to change your mind!

`

	fmt.Printf(msg, n, probNoSwitch, probSwitch)
}
