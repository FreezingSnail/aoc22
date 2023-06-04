package main

import (
	cal "aoc/day1"
	rps "aoc/day2"
	rucksack "aoc/day3"
	camp "aoc/day4"
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

func day4() {
	contains := camp.FindContainedAssignments(os.Args[1])
	fmt.Println("contained assignments: ", contains)
	overlaps := camp.FindOverlapedAssignments(os.Args[1])
	fmt.Println("contained assignments: ", overlaps)
}

func main() {
	day4()
}
