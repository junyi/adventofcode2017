package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt"

func computeSeverity(delay int, firewall [][]int) int {
	severity := 0
	for _, layer := range firewall {
		time := layer[0] + delay
		depth := layer[1]
		prevPos := time % (2*depth - 2)
		if prevPos == 0 {
			severity += time * depth
		}
	}

	return severity
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	firewall := make([][]int, 0)
	count := 0
	for scanner.Scan() {
		input := scanner.Text()
		arr := strings.Split(input, ": ")

		id, _ := strconv.Atoi(arr[0])
		depth, _ := strconv.Atoi(arr[1])

		layer := make([]int, 2)
		layer[0] = id
		layer[1] = depth

		firewall = append(firewall, layer)

		count++
	}

	delay := 0
	for {
		severity := computeSeverity(delay, firewall)

		if severity == 0 {
			break
		}

		delay++
	}

	fmt.Println("Result =", delay)
}
