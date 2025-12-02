package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input, err := os.ReadFile("ex.txt")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	p1 := 0
	p2 := 0

	inputstring := strings.SplitSeq(string(input), ",")
	for s := range inputstring {
		s = strings.Trim(s, "\n")
		r := strings.Split(s, "-")

		// Check if ID is valid
		if r[0][0] == '0' || r[1][0] == '0' {
			continue
		}

		rStart, err := strconv.Atoi(r[0])
		if err != nil {
			log.Fatal(err)
		}
		rEnd, err := strconv.Atoi(r[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := rStart; i <= rEnd; i++ {
			num := strconv.Itoa(i)
			num1 := num[:len(num)/2]
			num2 := num[len(num)/2:]
			if num1 == num2 {
				p1 += i
				p2 += i
			} else {
				for j := 0; j <= len(num)/2; j++ {
					cnt := strings.Count(num, num[:j])
					if cnt != 0 && len(num)%cnt == 0 && j*cnt == len(num) {
						p2 += i
						break
					}
				}
			}
		}
	}

	fmt.Println("P1=", p1)
	fmt.Println("P2=", p2)
}
