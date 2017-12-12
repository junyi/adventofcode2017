package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt"

func traverse(graph map[int]map[int]bool, visited map[int]bool, start int) {
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

	visited := make(map[int]bool)
	groups := 0

	for i := 0; i < len(graph); i++ {
		hasVisited := visited[i]
		if !hasVisited {
			traverse(graph, visited, i)
			groups++
		}
	}

	fmt.Println("Result =", groups)
}
