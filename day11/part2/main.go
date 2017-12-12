package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var path = "input.txt"

// An Item is something we manage in a priority queue.
type Item struct {
	x        float64
	y        float64
	steps    int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) get(i int) *Item {
	return (*pq)[i]
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func move(x float64, y float64, direction string) (float64, float64) {
	switch direction {
	case "n":
		y--
	case "s":
		y++
	case "ne":
		x += 0.5
		y -= 0.5
	case "se":
		x += 0.5
		y += 0.5
	case "nw":
		x -= 0.5
		y -= 0.5
	case "sw":
		x -= 0.5
		y += 0.5
	}

	return x, y
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func euclidean(x1, y1 float64, x2, y2 float64) int {
	x := math.Abs(x1 - x2)
	y := math.Abs(y1 - y2)
	return int(x*x*9 + y*y*3)
}

func distance(x float64, y float64) int {
	allDirs := []string{"n", "s", "ne", "se", "nw", "sw"}
	pq := &PriorityQueue{
		&Item{
			x:     0,
			y:     0,
			steps: 0,
			index: 0,
		},
	}
	heap.Init(pq)

	doneSet := make(map[string]bool)
	prevSet := make(map[string]int)
	doneSet["0.0,0.0"] = true

	var finalItem *Item

	for {
		if pq.Len() == 0 {
			break
		}

		i := heap.Pop(pq)
		item := i.(*Item)

		for _, dir := range allDirs {
			nx, ny := move(item.x, item.y, dir)
			priority := item.steps + 2 + euclidean(float64(x), float64(y), nx, ny)

			neighbour := &Item{
				x:        nx,
				y:        ny,
				steps:    item.steps + 2,
				priority: priority,
			}

			str := fmt.Sprintf("%.1f,%.1f", nx, ny)
			if _, ok := doneSet[str]; ok {
				continue
			} else {
				if index, ok := prevSet[str]; ok {
					if pq.get(index).priority > neighbour.priority {
						pq.update(pq.get(index), neighbour.priority)
					}
				} else {
					doneSet[str] = true
					neighbour.index = pq.Len() - 1
					heap.Push(pq, neighbour)
					prevSet[str] = neighbour.index
				}
			}
		}

		if item.x == x && item.y == y {
			finalItem = item
			break
		}
	}

	return int(finalItem.steps / 2)
}

func main() {
	// TODO: Not the best solution I have come up with
	// Should figure out a better one

	file, _ := ioutil.ReadFile(path)
	input := strings.Split(strings.Trim(string(file), "\n"), ",")

	var x, y float64
	var maxD int

	for _, dir := range input {
		x, y = move(x, y, dir)
		d := distance(x, y)
		if d > maxD {
			maxD = d
		}
	}

	fmt.Println("Result =", maxD)
}
