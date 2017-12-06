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
		prevNums := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			num, _ := strconv.Atoi(arr[i])
			for j := 0; j < i; j++ {
				prevNum := prevNums[j]
				if num%prevNum == 0 {
					sum += num / prevNum
				} else if prevNum%num == 0 {
					sum += prevNum / num
				}
			}
			prevNums[i] = num
		}
	}
	println("Result =", sum)
}
