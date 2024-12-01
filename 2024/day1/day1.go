package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

func convertDataToInt(sl [][]byte) ([]int64, []int64) {
	var leftData []int64
	var rightData []int64
	delimiter := []byte("   ")

	for _, slice := range sl {
		parts := bytes.Split(slice, delimiter)
		val, err := strconv.ParseInt(string(parts[0]), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		leftData = append(leftData, val)
		val, err = strconv.ParseInt(string(parts[1]), 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		rightData = append(rightData, val)
	}
	return leftData, rightData
}

func calcTotalDistance(d1 []int64, d2 []int64) int64 {
	var dist int64
	for i := 0; i < len(d1); i++ {
		dist += int64(math.Abs(float64(d1[i] - d2[i])))
	}
	return dist
}

func calcSimilarityScore(d1 []int64, d2 []int64) int64 {
	var score int64

	for _, lvalue := range d1 {
		var occurences int = 0

		for _, rvalue := range d2 {
			if lvalue == rvalue {
				occurences++
			}
		}
		score += lvalue * int64(occurences)
	}

	return score
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = bytes.Trim(data, "\n")
	sl := bytes.Split(data, []byte("\n"))

	d1, d2 := convertDataToInt(sl)

	slices.Sort(d1)
	slices.Sort(d2)

	p1 := calcTotalDistance(d1, d2)
	p2 := calcSimilarityScore(d1, d2)

	fmt.Println("Part1:", p1)
	fmt.Println("Part2:", p2)
}
