package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const path = "input.txt"

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Fields(line)
		min, _ := strconv.Atoi(arr[0])
		max := min
		for i := 1; i < len(arr); i++ {
			num, _ := strconv.Atoi(arr[i])
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		sum += (max - min)
	}
	println("Result =", sum)
}
