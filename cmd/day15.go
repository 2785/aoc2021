package cmd

import (
	"container/heap"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[15] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay15P1,
		P2: DoDay15P2,
	}
}

type d15QueueItem struct {
	point  day15Point
	weight int
	ind    int
}

type d15PQueue []*d15QueueItem

func (pq d15PQueue) Len() int { return len(pq) }

func (pq d15PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].ind = i
	pq[j].ind = j
}

func (pq d15PQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq *d15PQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*d15QueueItem)
	item.ind = n
	*pq = append(*pq, item)
}

func (pq *d15PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.ind = -1
	*pq = old[0 : n-1]
	return item
}

func DoDay15P1(input string) {
	sol, err := solveDay15P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 15 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

type day15Point struct{ x, y int }

func solveDay15P1(input string) (int, error) {
	lines := strings.Split(input, "\n")

	grid := make(map[day15Point]*struct {
		visited bool
		weight  int
	})

	for y, line := range lines {
		split := strings.Split(line, "")
		for x, c := range split {
			num, err := strconv.Atoi(c)
			if err != nil {
				return 0, err
			}

			grid[day15Point{x, y}] = &struct {
				visited bool
				weight  int
			}{
				visited: false,
				weight:  num,
			}
		}
	}

	queue := &d15PQueue{{point: day15Point{0, 0}, weight: 0, ind: 0}}
	heap.Init(queue)

	end := day15Point{len(lines[0]) - 1, len(lines) - 1}

	goal := 0

	for queue.Len() > 0 {
		currItem := heap.Pop(queue).(*d15QueueItem)
		if currItem.point == end {
			goal = currItem.weight
			break
		}

		currPoint := currItem.point

		grid[currPoint].visited = true

		left := day15Point{currPoint.x - 1, currPoint.y}
		right := day15Point{currPoint.x + 1, currPoint.y}
		up := day15Point{currPoint.x, currPoint.y - 1}
		down := day15Point{currPoint.x, currPoint.y + 1}

		if node, ok := grid[left]; ok && !node.visited {
			dist := currItem.weight + grid[left].weight
			heap.Push(queue, &d15QueueItem{point: left, weight: dist})
		}

		if node, ok := grid[right]; ok && !node.visited {
			dist := currItem.weight + grid[right].weight
			heap.Push(queue, &d15QueueItem{point: right, weight: dist})
		}

		if node, ok := grid[up]; ok && !node.visited {
			dist := currItem.weight + grid[up].weight
			heap.Push(queue, &d15QueueItem{point: up, weight: dist})
		}

		if node, ok := grid[down]; ok && !node.visited {
			dist := currItem.weight + grid[down].weight
			heap.Push(queue, &d15QueueItem{point: down, weight: dist})
		}
	}

	return goal, nil
}

func DoDay15P2(input string) {
	sol, err := solveDay15P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 15 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay15P2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	grid := make(map[day15Point]*struct {
		visited bool
		weight  int
	})

	for y, line := range lines {
		split := strings.Split(line, "")
		for x, c := range split {
			num, err := strconv.Atoi(c)
			if err != nil {
				return 0, err
			}

			grid[day15Point{x, y}] = &struct {
				visited bool
				weight  int
			}{
				visited: false,
				weight:  num,
			}
		}
	}

	dx, dy := len(lines[0]), len(lines)

	for sy := 1; sy <= 4; sy++ {
		for x := 0; x < dx; x++ {
			for y := 0; y < dy; y++ {
				p := day15Point{x, y + dy*sy}
				weight := grid[day15Point{x, y}].weight + sy
				if weight > 9 {
					weight = weight % 9
				}

				grid[p] = &struct {
					visited bool
					weight  int
				}{
					visited: false,
					weight:  weight,
				}

			}
		}
	}

	for sx := 1; sx <= 4; sx++ {
		for x := 0; x < dx; x++ {
			for y := 0; y < dy*5; y++ {
				p := day15Point{x + dx*sx, y}
				weight := grid[day15Point{x, y}].weight + sx
				if weight > 9 {
					weight = weight % 9
				}

				grid[p] = &struct {
					visited bool
					weight  int
				}{
					visited: false,
					weight:  weight,
				}

			}
		}
	}

	end := day15Point{len(lines[0])*5 - 1, len(lines)*5 - 1}

	queue := &d15PQueue{{point: day15Point{0, 0}, weight: 0, ind: 0}}
	heap.Init(queue)

	goal := 0

	for queue.Len() > 0 {
		currItem := heap.Pop(queue).(*d15QueueItem)
		if currItem.point == end {
			goal = currItem.weight
			break
		}

		currPoint := currItem.point

		grid[currPoint].visited = true

		left := day15Point{currPoint.x - 1, currPoint.y}
		right := day15Point{currPoint.x + 1, currPoint.y}
		up := day15Point{currPoint.x, currPoint.y - 1}
		down := day15Point{currPoint.x, currPoint.y + 1}

		if node, ok := grid[left]; ok && !node.visited {
			dist := currItem.weight + grid[left].weight
			heap.Push(queue, &d15QueueItem{point: left, weight: dist})
		}

		if node, ok := grid[right]; ok && !node.visited {
			dist := currItem.weight + grid[right].weight
			heap.Push(queue, &d15QueueItem{point: right, weight: dist})
		}

		if node, ok := grid[up]; ok && !node.visited {
			dist := currItem.weight + grid[up].weight
			heap.Push(queue, &d15QueueItem{point: up, weight: dist})
		}

		if node, ok := grid[down]; ok && !node.visited {
			dist := currItem.weight + grid[down].weight
			heap.Push(queue, &d15QueueItem{point: down, weight: dist})
		}
	}
	return goal, nil

}
