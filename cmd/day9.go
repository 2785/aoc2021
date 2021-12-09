package cmd

import (
	"errors"
	"sort"
	"strings"

	"github.com/2785/aoc2021/util"
	"go.uber.org/zap"
)

func init() {
	solvers[9] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay9P1,
		P2: DoDay9P2,
	}
}

func DoDay9P1(input string) {
	sol, err := solveDay9P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 9 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay9P1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	tile := make([][]int, len(lines))
	for i, line := range lines {
		row := strings.Split(line, "")
		rowNum, err := util.IntsFromStrings(row)
		if err != nil {
			return 0, err
		}

		tile[i] = rowNum
	}

	score := 0

	for i := 0; i < len(tile); i++ {
		for j := 0; j < len(tile[i]); j++ {
			adj := make([]int, 0)
			if i-1 >= 0 {
				adj = append(adj, tile[i-1][j])
			}

			if i+1 < len(tile) {
				adj = append(adj, tile[i+1][j])
			}

			if j-1 >= 0 {
				adj = append(adj, tile[i][j-1])
			}

			if j+1 < len(tile[i]) {
				adj = append(adj, tile[i][j+1])
			}

			lowPoint := true
			for _, v := range adj {
				if v <= tile[i][j] {
					lowPoint = false
				}
			}

			if lowPoint {
				score += tile[i][j] + 1
			}
		}
	}

	return score, nil
}

func DoDay9P2(input string) {
	sol, err := solveDay9P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 9 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

type day9Point struct {
	x, y int
}

func solveDay9P2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	tile := make([][]int, len(lines))
	for i, line := range lines {
		row := strings.Split(line, "")
		rowNum, err := util.IntsFromStrings(row)
		if err != nil {
			return 0, err
		}

		tile[i] = rowNum
	}

	lowPoints := []day9Point{}

	for i := 0; i < len(tile); i++ {
		for j := 0; j < len(tile[i]); j++ {
			adj := make([]int, 0)
			if i-1 >= 0 {
				adj = append(adj, tile[i-1][j])
			}

			if i+1 < len(tile) {
				adj = append(adj, tile[i+1][j])
			}

			if j-1 >= 0 {
				adj = append(adj, tile[i][j-1])
			}

			if j+1 < len(tile[i]) {
				adj = append(adj, tile[i][j+1])
			}

			lowPoint := true
			for _, v := range adj {
				if v <= tile[i][j] {
					lowPoint = false
				}
			}

			if lowPoint {
				lowPoints = append(lowPoints, day9Point{x: j, y: i})
			}
		}
	}

	basinSizes := []int{}

	for _, lp := range lowPoints {
		size := findBasinSize(tile, lp)

		basinSizes = append(basinSizes, size)
	}

	if len(basinSizes) < 3 {
		return 0, errors.New("wtf")
	}

	sort.Ints(basinSizes)

	prod := 1

	for i := 0; i < 3; i++ {
		prod *= basinSizes[len(basinSizes)-1-i]
	}

	return prod, nil
}

func findBasinSize(grid [][]int, point day9Point) int {
	pointMap := make(map[day9Point]bool)
	width := len(grid[0])
	height := len(grid)

	pointMap[point] = true

	neighbors := []day9Point{point}

	doneOriginal := false

	attempt := 0

	for attempt <= 500 {
		attempt++
		current := make([]day9Point, len(neighbors))
		copy(current, neighbors)

		for _, n := range current {
			if pointMap[n] && (n != point || doneOriginal) {
				continue
			}
			doneOriginal = true
			subNeighbors := []day9Point{}
			if n.x-1 >= 0 && !pointMap[day9Point{x: n.x - 1, y: n.y}] && grid[n.y][n.x-1] < 9 {
				subNeighbors = append(subNeighbors, day9Point{x: n.x - 1, y: n.y})
			}

			if n.x+1 < width && !pointMap[day9Point{x: n.x + 1, y: n.y}] && grid[n.y][n.x+1] < 9 {
				subNeighbors = append(subNeighbors, day9Point{x: n.x + 1, y: n.y})
			}

			if n.y-1 >= 0 && !pointMap[day9Point{x: n.x, y: n.y - 1}] && grid[n.y-1][n.x] < 9 {
				subNeighbors = append(subNeighbors, day9Point{x: n.x, y: n.y - 1})
			}

			if n.y+1 < height && !pointMap[day9Point{x: n.x, y: n.y + 1}] && grid[n.y+1][n.x] < 9 {
				subNeighbors = append(subNeighbors, day9Point{x: n.x, y: n.y + 1})
			}

			values := []int{}
			for _, v := range subNeighbors {
				values = append(values, grid[v.y][v.x])
			}

			low := true
			for _, v := range values {
				if v < grid[n.y][n.x] {
					low = false
				}
			}

			if low {
				pointMap[n] = true
				for _, sub := range subNeighbors {
					if !inPoints(neighbors, sub) {
						neighbors = append(neighbors, sub)
					}
				}
			}
		}
	}

	return len(neighbors)
}

func inPoints(coll []day9Point, p day9Point) bool {
	for _, v := range coll {
		if v == p {
			return true
		}
	}

	return false
}
