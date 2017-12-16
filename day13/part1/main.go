package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt"

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

	severity := 0
	for _, layer := range firewall {
		time := layer[0]
		depth := layer[1]
		prevPos := time % (2*depth - 2)
		if prevPos == 0 {
			severity += time * depth
		}
	}

	fmt.Println("Result =", severity)
}
