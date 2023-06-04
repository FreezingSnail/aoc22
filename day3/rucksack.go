package rucksack

import (
	"aoc/reader"
)

type compartment struct {
	items map[rune]interface{}
}
type sack struct {
	as string
	a  compartment
	bs string
	b  compartment
}
type group struct {
	sacks [3]sack
}

func stringToCompartment(str string) compartment {
	items := make(map[rune]interface{}, 0)
	cmpt := compartment{items: items}
	for _, c := range str {
		if _, ok := cmpt.items[c]; !ok {
			cmpt.items[c] = nil
		}
	}

	return cmpt
}

func findOverlap(a, b compartment) rune {
	for key, _ := range a.items {
		if _, ok := b.items[key]; ok {
			//fmt.Println("matched key:", key, string(key))
			return key
		}
	}
	return -1
}

func priority(c rune) int {
	if c <= 'z' && c >= 'a' {
		return int(byte(c) - ('a' - 1))
	}
	return int(byte(c) - ('A' - 1) + 26)

}

func stringToSack(s string) sack {
	size := len(s) / 2
	saq := sack{
		as: s[:size],
		a:  stringToCompartment(s[:size]),
		bs: s[size:],
		b:  stringToCompartment(s[size:]),
	}
	return saq
}

func (s sack) findPriority() int {
	return priority(findOverlap(s.a, s.b))
}

func listToSacks(list []string) []sack {
	sacks := make([]sack, len(list))
	for i, s := range list {
		sacks[i] = stringToSack(s)
	}
	return sacks
}

func SumPrioritys(dir string) int {
	_, list := reader.BuildCalList(dir)
	sacks := listToSacks(list)
	sum := 0
	for _, s := range sacks {
		//fmt.Println(s.as, s.bs)
		p := s.findPriority()
		//fmt.Println(p)
		sum += p
	}

	return sum

}

func (g group) groupOverlap() rune {
	//fmt.Println("group", g)
	keyCounts := make(map[rune]int, 0)
	for _, sack := range g.sacks {
		for k, _ := range sack.a.items {
			if _, ok := keyCounts[k]; ok {
				keyCounts[k] += 1
				if count, _ := keyCounts[k]; count == 3 {
					//fmt.Println("badge:", string(k), k)
					return k
				}
			} else {
				keyCounts[k] = 1
			}
		}
		for k, _ := range sack.b.items {
			if _, ok := keyCounts[k]; ok {
				keyCounts[k] += 1
				if count, _ := keyCounts[k]; count == 3 {
					//fmt.Println("badge:", string(k), k)
					return k
				}
			} else {
				keyCounts[k] = 1
			}
		}
	}
	//fmt.Println(keyCounts)
	return '0'
}

func buildGroups(sacks []sack) []group {
	groups := make([]group, 0)
	i := 0
	for i < len(sacks) {
		g := group{sacks: [3]sack{sacks[i], sacks[i+1], sacks[i+2]}}
		groups = append(groups, g)
		i += 3
	}
	//fmt.Println(len(groups))
	return groups
}

func SumGroupPriorities(dir string) int {
	_, list := reader.BuildCalList(dir)
	//fmt.Println(len(list))
	sacks := listToSacks(list)
	groups := buildGroups(sacks)
	sum := 0
	for _, g := range groups {
		p := g.groupOverlap()
		sum += priority(p)
	}

	return sum

}
