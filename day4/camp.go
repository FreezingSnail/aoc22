package camp

import (
	"aoc/reader"
	"strconv"
	"strings"
)

type assignment struct {
	start int
	end   int
}

type pair struct {
	assignments [2]assignment
}

func toAssignment(s string) assignment {
	splits := strings.Split(s, string("-"))
	start, err := strconv.Atoi(splits[0])
	if err != nil {
		return assignment{}
	}
	end, err := strconv.Atoi(splits[1])
	if err != nil {
		return assignment{}
	}
	return assignment{start: start, end: end}
}

func stringToPair(s string) pair {
	splits := strings.Split(s, ",")
	return pair{assignments: [2]assignment{
		toAssignment(splits[0]),
		toAssignment(splits[1]),
	}}
}

func listToPairs(list []string) []pair {
	pairs := make([]pair, len(list))
	for i, s := range list {
		pairs[i] = stringToPair(s)
	}

	return pairs
}

func (p pair) containAssignment() bool {
	if p.assignments[0].start >= p.assignments[1].start &&
		p.assignments[0].end <= p.assignments[1].end {
		return true
	}
	if p.assignments[1].start >= p.assignments[0].start &&
		p.assignments[1].end <= p.assignments[0].end {
		return true
	}

	return false
}

func (p pair) overlaps() bool {
	if p.assignments[0].start >= p.assignments[1].start &&
		p.assignments[0].start <= p.assignments[1].end {
		return true
	}
	if p.assignments[0].end >= p.assignments[1].start &&
		p.assignments[0].end <= p.assignments[1].end {
		return true
	}

	if p.assignments[1].end >= p.assignments[0].start &&
		p.assignments[1].end <= p.assignments[0].end {
		return true
	}
	if p.assignments[1].start >= p.assignments[0].start &&
		p.assignments[1].start <= p.assignments[0].end {
		return true
	}

	return false
}

func FindContainedAssignments(dir string) int {
	_, list := reader.BuildCalList(dir)
	pairs := listToPairs(list)
	count := 0
	for _, p := range pairs {
		if p.containAssignment() {
			count += 1
		}
	}
	return count
}

func FindOverlapedAssignments(dir string) int {
	_, list := reader.BuildCalList(dir)
	pairs := listToPairs(list)
	count := 0
	for _, p := range pairs {
		if p.overlaps() {
			count += 1
		}
	}
	return count
}
