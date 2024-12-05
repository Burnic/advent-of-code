package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(rulemap map[string][]string, updates []string) int {
	var sum int = 0

	for _, update := range updates {
		ups := strings.Split(update, ",")
		var isValid bool = true
		for idx, up := range ups {
			for i := 0; i < idx; i++ {
				if slices.Contains(rulemap[up], ups[i]) {
					isValid = false
					break
				}
			}
		}
		if isValid {
			midIdx := (len(ups) - 1) / 2
			val, err := strconv.Atoi(ups[midIdx])
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}
	}
	return sum
}
func part2(rulemap map[string][]string, updates []string) int {
	var sum int = 0

	for _, update := range updates {
		ups := strings.Split(update, ",")
		var isChanged bool = false

	reset:
		for idx, up := range ups {
			for i := 0; i < idx; i++ {
				if slices.Contains(rulemap[up], ups[i]) {
					ups[idx], ups[idx-1] = ups[idx-1], ups[idx]
					isChanged = true
					goto reset
				}
			}
		}
		if isChanged {
			midIdx := (len(ups) - 1) / 2
			val, err := strconv.Atoi(ups[midIdx])
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = bytes.Trim(data, "\n")
	s := string(data)

	sl := strings.Split(s, "\n")
	idx := slices.Index(sl, "")
	if idx == -1 {
		log.Fatal("Seperator not found")
	}

	rules := sl[:idx]
	updates := sl[idx+1:]

	rulemap := make(map[string][]string)

	for _, rule := range rules {
		values := strings.Split(rule, "|")
		rulemap[values[0]] = append(rulemap[values[0]], values[1])
	}
	p1 := part1(rulemap, updates)
	p2 := part2(rulemap, updates)

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}
