package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertDataToInt(sl [][]byte) [][]byte {
	var output [][]byte

	for _, row := range sl {
		var byteRow []byte
		data := bytes.Split(row, []byte(" "))
		for _, val := range data {
			intVal, err := strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
			byteRow = append(byteRow, byte(intVal))
		}
		output = append(output, byteRow)
	}
	return output
}

func findNumSafeReports(reports [][]byte) int {
	var numSafeReports int = 0

	for _, report := range reports {
		var prevLevel byte = 0
		if report[0] < report[1] {
			for i, level := range report {
				if i != 0 {
					if level-prevLevel > 3 || level-prevLevel < 1 {
						break
					}
					if i == len(report)-1 {
						numSafeReports++
					}
				}
				prevLevel = level
			}
		} else {
			for i, level := range report {
				if i != 0 {
					if prevLevel-level > 3 || prevLevel-level < 1 {
						break
					}
					if i == len(report)-1 {
						numSafeReports++
					}
				}
				prevLevel = level
			}
		}
	}

	return numSafeReports
}
func findNumSafeReportsDampen(reports [][]byte) int {
	var numSafeReports int = 0

	for _, report := range reports {
		var prevLevel byte = 0
		var dampen bool = false
		var dir int = 0
		var tryAgain bool = false

		for i, level := range report {
			if i == 0 {
				if level > report[len(report)-1] {
					dir = -1
				} else if level < report[len(report)-1] {
					dir = 1
				} else {
					break
				}
			}
			if i != 0 {
				if dir == 1 {
					if level-prevLevel > 3 || level-prevLevel < 1 {
						if !dampen {
							dampen = true
							if i == 1 {
								tryAgain = true
							}
							if i == len(report)-1 {
								numSafeReports++
							}
							continue
						}
						if i == 2 && tryAgain {
							if level-report[1] > 3 || level-report[1] < 1 {
							} else {
								prevLevel = level
								continue
							}
						}
						break
					}
				} else if dir == -1 {
					if prevLevel-level > 3 || prevLevel-level < 1 {
						if !dampen {
							dampen = true
							if i == 1 {
								tryAgain = true
							}
							if i == len(report)-1 {
								numSafeReports++
							}
							continue
						}
						if i == 2 && tryAgain {
							if report[1]-level > 3 || report[1]-level < 1 {
							} else {
								prevLevel = level
								continue
							}
						}
						break
					}
				}
				if i == len(report)-1 {
					numSafeReports++
				}
			}
			prevLevel = level
		}
	}
	return numSafeReports
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data = bytes.Trim(data, "\n")
	sl := bytes.Split(data, []byte("\n"))

	d := convertDataToInt(sl)
	p1 := findNumSafeReports(d)
	fmt.Println("p1:", p1)

	p2 := findNumSafeReportsDampen(d)
	fmt.Println("p2:", p2)
}
