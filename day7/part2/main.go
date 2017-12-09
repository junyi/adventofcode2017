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
	Name           string
	Weight         int
	Children       []string
	ChildrenWeight []int
}

func allEqual(arr []int) bool {
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] != prev {
			return false
		}
		prev = arr[i]
	}
	return true
}

func findOneDifferent(arr []int) int {
	length := len(arr)

	firstDiffIdx := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[i-1] {
			firstDiffIdx = i
			break
		}
	}

	if arr[firstDiffIdx] == arr[(firstDiffIdx-2+length)%length] {
		return (firstDiffIdx - 1 + length) % length
	}
	return firstDiffIdx
}

func checkBalanced(node string, programMap map[string]Program) (int, bool, string) {
	program := programMap[node]
	children := program.Children
	weight := program.Weight
	cWeights := program.ChildrenWeight

	if children == nil {
		return weight, true, node
	}

	for i, child := range children {
		w, b, n := checkBalanced(child, programMap)
		cWeights[i] = w
		weight += w

		if !b {
			return w, false, n
		}
	}

	return weight, allEqual(cWeights), node
}

func findRoot(childToParent map[string]string) string {
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

	return child
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
			Name:           name,
			Weight:         weight,
			Children:       children,
			ChildrenWeight: make([]int, len(children)),
		}

		programMap[name] = program
	}

	root := findRoot(childToParent)

	_, _, n := checkBalanced(root, programMap)
	unbalancedProgram := programMap[n]
	children := unbalancedProgram.Children
	childrenWeight := unbalancedProgram.ChildrenWeight

	wrongIdx := findOneDifferent(unbalancedProgram.ChildrenWeight)
	correctIdx := (wrongIdx + 1) % len(unbalancedProgram.Children)

	diffInWeight := childrenWeight[wrongIdx] - childrenWeight[correctIdx]

	result := programMap[children[wrongIdx]].Weight - diffInWeight

	fmt.Println("Result =", result)
}
