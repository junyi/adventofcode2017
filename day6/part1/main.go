package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const path = "input.txt"

func distribute(arr []int) {
	var i int
	l := len(arr)
	argmax := 0
	max := arr[0]
	for i = 1; i < l; i++ {
		if arr[i] > max {
			argmax = i
			max = arr[i]
		}
	}

	arr[argmax] = 0

	quotient := max / l
	remainder := max % l
	for j := 0; j < l; j++ {
		i = (argmax + 1 + j) % l
		arr[i] += quotient
		if remainder > 0 {
			arr[i]++
			remainder--
		}
	}
}

func config(arr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	steps := 0
	scanner.Scan()
	line := scanner.Text()
	strArr := strings.Fields(line)
	arr := make([]int, len(strArr))
	uniqueMap := make(map[string]bool)
	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

	for {
		c := config(arr)
		if _, contains := uniqueMap[c]; contains {
			break
		} else {
			uniqueMap[c] = true
		}
		steps++
		distribute(arr)
	}

	println("Result =", steps)
}
