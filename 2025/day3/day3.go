package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findNextIndex(s string) int {
	idx := -1

	for i := '9'; i >= '0'; i-- {
		idx = strings.IndexByte(s, byte(i))
		if idx != -1 {
			break
		}
	}
	return idx
}

func findBat(numBat int, s string) int {
	startIdx := 0
	bat := ""
	x := make([]int, 0, numBat)

	for i := numBat - 1; len(x) < numBat; i-- {
		idx := findNextIndex(s[startIdx : len(s)-i])
		x = append(x, idx+startIdx)
		startIdx = x[len(x)-1] + 1
		bat += string(s[x[len(x)-1]])
	}

	num, err := strconv.Atoi(bat)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	p1, p2 := 0, 0

	for _, s := range strings.Fields(string(input)) {
		p1 += findBat(2, s)
		p2 += findBat(12, s)
	}

	fmt.Println("P1=", p1)
	fmt.Println("P2=", p2)
}
