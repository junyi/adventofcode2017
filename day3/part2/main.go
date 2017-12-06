package main

const input = 265149

func set(dict map[int]map[int]int, x int, y int, val int) {
	if _, ok := dict[x]; !ok {
		dict[x] = make(map[int]int)
	}
	dict[x][y] = val
}

func get(dict map[int]map[int]int, x int, y int) int {
	if inner, ok := dict[x]; ok {
		if val, ok2 := inner[y]; ok2 {
			return val
		}
	}
	return 0
}

func sumNeighbours(dict map[int]map[int]int, x int, y int) int {
	neighbours := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	sum := 0

	for _, coord := range neighbours {
		s := get(dict, coord[0]+x, coord[1]+y)
		// println(">>>", coord[0]+x, coord[1]+y, s)
		sum += s
	}

	return sum
}

func main() {
	direction := [4][2]int{
		{1, 0},  // right
		{0, -1}, // up
		{-1, 0}, // left
		{0, 1},  // down
	}

	spiral := make(map[int]map[int]int)
	x := 0
	y := 0
	num := 1
	multiple := 1
	dirIndex := 0
	counter := 0
	innerCounter := 0
	set(spiral, x, y, num)

	for {
		if num > 1 {
			sum := sumNeighbours(spiral, x, y)
			set(spiral, x, y, sum)
			if sum > input {
				println("Result =", sum)
				break
			}
			// println(x, y, num, sum)
		}

		dir := direction[dirIndex]
		x += dir[0]
		y += dir[1]
		num++
		innerCounter++
		if innerCounter == multiple {
			counter = (counter + 1) % 2
			dirIndex = (dirIndex + 1) % 4
			innerCounter = 0
			if counter == 0 {
				multiple++
			}
		}
	}
}
