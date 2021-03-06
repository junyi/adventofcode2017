package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt"

func traverse(graph map[int]map[int]bool, start int) int {
	visited := make(map[int]bool)
	queue := []int{start}

	for {
		if len(queue) == 0 {
			break
		}

		item := queue[0]
		queue = queue[1:]

		visited[item] = true

		for neighbour := range graph[item] {
			if _, ok := visited[neighbour]; !ok {
				queue = append(queue, neighbour)
			}
		}
	}

	return len(visited)
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	graph := make(map[int]map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " <-> ")
		id, _ := strconv.Atoi(arr[0])
		tempArr := strings.Split(arr[1], ", ")

		for _, v := range tempArr {
			n, _ := strconv.Atoi(v)
			if _, ok := graph[id]; !ok {
				graph[id] = make(map[int]bool)
			}
			if _, ok := graph[n]; !ok {
				graph[n] = make(map[int]bool)
			}

			graph[id][n] = true
			graph[n][id] = true
		}
	}

	length := traverse(graph, 0)

	fmt.Println("Result =", length)
}
