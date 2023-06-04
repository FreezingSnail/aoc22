package calorie

import (
	"aoc/reader"
	"sort"
	"strconv"
)

type elf struct {
	id       int
	calories int
}

func MostCalories(dir string) (error, int) {
	err, list := reader.BuildCalList(dir)
	if err != nil {
		return err, 0
	}

	err, elfs := buildElfList(list)
	if err != nil {
		return err, 0
	}

	return nil, elfs[0].calories
}

func TopThreeCalories(dir string) (error, int) {
	err, list := reader.BuildCalList(dir)
	if err != nil {
		return err, 0
	}

	err, elfs := buildElfList(list)
	if err != nil {
		return err, 0
	}

	return nil, elfs[0].calories + elfs[1].calories + elfs[2].calories
}

func buildElfList(calorieList []string) (error, []elf) {
	runningSum := 0
	curElf := 1
	var elfs []elf

	for _, c := range calorieList {
		if c == "" {
			elfs = append(elfs, elf{curElf, runningSum})
			runningSum = 0
			curElf += 1
			continue
		}
		cur, err := strconv.Atoi(c)
		if err != nil {
			return err, nil
		}

		runningSum += cur
	}

	sort.Slice(elfs[:], func(i, j int) bool {
		return elfs[i].calories > elfs[j].calories
	})

	return nil, elfs
}
