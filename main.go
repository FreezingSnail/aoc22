package main

import (
	"fmt"
	cal "aoc/day1"
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


func main() {
}
