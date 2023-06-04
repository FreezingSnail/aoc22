package rps

import (
	"aoc/reader"
	"fmt"
)

const x int = 1
const y int = 2
const z int = 3

var lookup = map[string]int{"A": 1, "B": 2, "C": 3, "X": x, "Y": y, "Z": z}

type game struct {
	opp string
	plr string
}

func (g game) outcome() int {
	switch lookup[g.opp] {
	case 1: // rock
		switch lookup[g.plr] {
		case 1: //rock
			return 3
		case 2: //paper
			return 6
		case 3: //scissors
			return 0
		}
	case 2: // paper
		switch lookup[g.plr] {
		case 1: //rock
			return 0
		case 2: //paper
			return 3
		case 3: //scissors
			return 6
		}
	case 3: // scissors
		switch lookup[g.plr] {
		case 1: //rock
			return 6
		case 2: //paper
			return 0
		case 3: //scissors
			return 3
		}
	}

	return -1
}

func (g game) score() int {
	return lookup[g.plr] + g.outcome()
}

func listToGames(list []string) []game {
	var games []game
	for _, c := range list {
		games = append(games, game{string(c[0]), string(c[2])})
	}
	return games
}

func calcScore(games []game) int {
	score := 0
	for _, g := range games {
		score += g.score()
	}
	return score
}

func RpsScore(dir string) int {
	err, list := reader.BuildCalList(dir)
	if err != nil {
		panic(err)
	}

	games := listToGames(list)
	return calcScore(games)

}

// PART 2

func (g game) outcome2() int {
	switch g.plr {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	fmt.Println("outcome fail")
	return -1
}

func (g game) decodePlayer() int {
	switch lookup[g.opp] {
	case 1: // rock
		switch g.plr {
		case "X": //loose
			return 3
		case "Y": //draw
			return 1
		case "Z": //win
			return 2
		}
	case 2: // paper
		switch g.plr {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		}

	case 3: // scissors
		switch g.plr {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		}
	}
	fmt.Println("decode fail")
	return -1
}

func (g game) score2() int {
	score := g.decodePlayer() + g.outcome2()
	return score
}

func calcScore2(games []game) int {
	score := 0
	for _, g := range games {
		score += g.score2()
	}
	return score
}

func RpsScore2(dir string) int {
	err, list := reader.BuildCalList(dir)
	if err != nil {
		panic(err)
	}

	games := listToGames(list)
	return calcScore2(games)

}
