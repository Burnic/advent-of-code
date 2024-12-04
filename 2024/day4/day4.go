package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const Keyword string = "XMAS"
const KeywordReversed string = "SAMX"

func getWordCount(input string) (int, int) {
	var p1count int = 0
	var p2count int = 0

	lines := strings.Split(input, "\n")

	for lineNum, line := range lines {

		p1count += findKeyword(line, lines[lineNum:], Keyword)
		p1count += findKeyword(line, lines[lineNum:], KeywordReversed)

		if lineNum != 0 && lineNum != len(lines)-1 {
			p2count += findXmas(line, lines[lineNum-1:lineNum+2])
		}
	}
	return p1count, p2count
}

func findKeyword(line string, lines []string, s string) int {
	var count int = 0
	startIndex := strings.Index(line, string(s[0]))

	for startIndex >= 0 {
		count += checkHorizontal(line[startIndex:], s)
		count += checkVertical(lines, startIndex, s)
		count += checkSlash(lines, startIndex, s)
		count += checkBackSlash(lines, startIndex, s)

		nextIndex := strings.Index(line[startIndex+1:], string(s[0]))
		if nextIndex == -1 {
			break
		}
		startIndex += nextIndex + 1
	}
	return count
}

func findXmas(line string, lines []string) int {
	var count int = 0
	startIndex := strings.Index(line, "A")

	for startIndex >= 0 {
		if startIndex != 0 {
			count += checkXmas(lines, startIndex)
		}
		nextIndex := strings.Index(line[startIndex+1:], "A")
		if nextIndex == -1 {
			break
		}
		startIndex += nextIndex + 1
	}
	return count
}

func checkXmas(lines []string, idx int) int {
	if len(lines) != 3 || idx > len(lines[0])-2 {
		return 0
	}
	s1 := string(lines[0][idx-1])
	s1 += string(lines[1][idx])
	s1 += string(lines[2][idx+1])

	s2 := string(lines[0][idx+1])
	s2 += string(lines[1][idx])
	s2 += string(lines[2][idx-1])

	if (s1 == "MAS" || s1 == "SAM") && (s2 == "MAS" || s2 == "SAM") {
		return 1
	}

	return 0
}

func checkHorizontal(s string, target string) int {
	if len(s) < len(target) {
		return 0
	} else if s[:len(target)] == target {
		return 1
	} else {
		return 0
	}
}

func checkVertical(sl []string, idx int, target string) int {
	if len(sl) < len(target) {
		return 0
	}
	for i := 0; i < len(target); i++ {
		if sl[i][idx] != target[i] {
			return 0
		}
	}
	return 1
}

func checkSlash(sl []string, idx int, target string) int {
	if len(sl) < len(target) {
		return 0
	} else if len(sl[0][idx:]) < len(target) {
		return 0
	}
	for i := 0; i < len(target); i++ {
		if sl[i][idx+i] != target[i] {
			return 0
		}
	}
	return 1
}

func checkBackSlash(sl []string, idx int, target string) int {
	if len(sl) < len(target) {
		return 0
	} else if idx < len(target)-1 {
		return 0
	}
	for i := 0; i < len(target); i++ {
		if sl[i][idx-i] != target[i] {
			return 0
		}
	}
	return 1
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = bytes.Trim(data, "\n")
	s := string(data)

	p1, p2 := getWordCount(s)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)

}
