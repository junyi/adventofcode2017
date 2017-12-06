package main

import (
	"bufio"
	"os"
	"strconv"
)

const path = "input.txt"

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	arr := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		arr = append(arr, num)
	}

	i := 0
	steps := 0
	for {
		if i < 0 || i >= len(arr) {
			break
		}
		steps++
		offset := arr[i]
		if offset >= 3 {
			arr[i]--
		} else {
			arr[i]++
		}
		i += offset
	}
	println("Result =", steps)
}
