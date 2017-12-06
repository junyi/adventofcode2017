package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

const path = "input.txt"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

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
			sortedWord := SortString(word)
			if _, contains := uniqueMap[sortedWord]; contains {
				valid = false
				break
			} else {
				uniqueMap[sortedWord] = true
			}
		}

		if valid {
			count++
		}
	}
	println("Result =", count)
}
