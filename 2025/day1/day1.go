package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	dial := 50
	dir := 'R'

	p1, p2 := 0, 0
	for _, s := range strings.Fields(string(input)) {
		n, _ := strconv.Atoi(s[1:])

		if dir != rune(s[0]) {
			dial = (100 - dial) % 100
			dir = rune(s[0])
		}
		dial += n
		p2 += dial / 100
		dial %= 100

		if dial%100 == 0 {
			p1++
		}
	}
	fmt.Println("P1=", p1)
	fmt.Println("P2=", p2)
}
