package main

import "math"

const num = 265149

/*
	17  16  15  14  13
	18   5   4   3  12
	19   6   1   2  11
	20   7   8   9  10
	21  22  23  24  25

	Coordinates for squares:
		Even squares (2k)^2:
			(-k + 1, k)
		E.g. 16 (2 * 2)^2 is at (-1, 2)

		Odd squares (2k + 1)^2:
			(k, -k)
		E.g. 25 (2 * 2 + 1)^2 is at (2, 2)
*/

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	l := int(math.Sqrt(num))
	u := l + 1
	lower := l * l
	upper := u * u
	corner := (lower + upper + 1) / 2

	var result int
	var k int
	if l%2 == 0 {
		k = l / 2
	} else {
		k = u / 2
	}

	if num == lower && lower%2 == 0 {
		result = 2*k - 1
	} else {
		result = abs(abs(corner-num)-k) + k
	}

	println("Result =", result, k, corner, num, lower, upper)
}
