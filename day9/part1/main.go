package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "input.txt"

func Score(input string) int {
	groupScore := 0
	groups := 0
	garbageMode := false
	ignore := false
	for _, c := range input {
		if ignore {
			ignore = false
			continue
		}

		if !garbageMode {
			switch c {
			case '{':
				groups++
				groupScore += groups
			case '}':
				groups--
			case '!':
				ignore = true
			case '<':
				garbageMode = true
			}
		} else {
			switch c {
			case '!':
				ignore = true
			case '>':
				garbageMode = false
			}
		}
	}

	return groupScore
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		groupScore := Score(line)
		fmt.Println("Result =", groupScore)
	}

}
