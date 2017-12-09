package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const path = "input.txt"

type Program struct {
	Name     string
	Weight   int
	Children []string
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`([a-z]+) \((\d+)\)(?: -> (.+))?`)

	programMap := make(map[string]Program)
	childToParent := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		name := matches[1]
		weight, _ := strconv.Atoi(matches[2])
		var children []string
		if matches[3] != "" {
			children = strings.Split(matches[3], ", ")
			for _, child := range children {
				childToParent[child] = name
			}
		}

		program := Program{
			Name:     name,
			Weight:   weight,
			Children: children,
		}

		programMap[name] = program
	}

	var child, parent string
	var ok bool

	for _, parent := range childToParent {
		child = parent
		break
	}

	for {
		if parent, ok = childToParent[child]; !ok {
			break
		}
		child = parent
	}

	fmt.Println("Result =", child)
}
