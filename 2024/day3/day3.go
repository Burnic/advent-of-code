package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func multSum(input string, part2 bool) uint64 {
	var sum uint64 = 0

	inslice := strings.Split(input, "don't()")

	for i, sl := range inslice {
		if i != 0 && part2 {
			doIndex := strings.Index(sl, "do()")
			if doIndex >= 0 {
				sl = sl[doIndex:]
			} else {
				continue
			}
		}
		idx := strings.Index(sl, "mul(")
		sl = sl[idx:]
		mulSlice := strings.Split(sl, "mul(")

		for _, muls := range mulSlice {
			commaIndex := strings.Index(muls, ",")
			parenIndex := strings.Index(muls, ")")

			//each number can be max 3 digits
			validMul := commaIndex > 0 && commaIndex < 4 && parenIndex > commaIndex && parenIndex < 8
			if validMul {
				mulStr := muls[:parenIndex]
				mulSli := strings.Split(mulStr, ",")

				val1, err := strconv.Atoi(mulSli[0])
				if err != nil {
					log.Fatal(err)
				}
				val2, err := strconv.Atoi(mulSli[1])
				if err != nil {
					log.Fatal(err)
				}
				sum += uint64(val1 * val2)
			}
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
	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))
	s := string(data)

	p1 := multSum(s, false)
	p2 := multSum(s, true)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}
