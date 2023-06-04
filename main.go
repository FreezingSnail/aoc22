package main

import (
	cal "aoc/day1"
	rps "aoc/day2"
	rucksack "aoc/day3"
	"fmt"
	"os"
)

func day1() {
	err, cals := cal.MostCalories(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Most Calories: ", cals)
	}
	err, cals = cal.TopThreeCalories(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top 3 sum Calories: ", cals)
	}
}

func day2() {
	score := rps.RpsScore(os.Args[1])
	fmt.Println("score: ", score)
	score = rps.RpsScore2(os.Args[1])
	fmt.Println("score: ", score)
}

func day3() {
	sumPriority := rucksack.SumPrioritys(os.Args[1])
	fmt.Println("Priorities: ", sumPriority)
	groupPriorities := rucksack.SumGroupPriorities(os.Args[1])
	fmt.Println("Group Priorities: ", groupPriorities)
}

func main() {
	day3()
}
