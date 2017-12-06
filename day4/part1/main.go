package main

import (
	"bufio"
	"os"
	"strings"
)

const path = "input.txt"

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Fields(line)
		uniqueMap := make(map[string]bool)
		valid := true

		for i := 0; i < len(arr); i++ {
			word := arr[i]
			if _, contains := uniqueMap[word]; contains {
				valid = false
				break
			} else {
				uniqueMap[word] = true
			}
		}

		if valid {
			count++
		}
	}
	println("Result =", count)
}
