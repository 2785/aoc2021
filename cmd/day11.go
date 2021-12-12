package cmd

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[11] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay11P1,
		P2: DoDay11P2,
	}
}

func parseGrind(input string) (day11Grid, error) {
	lines := strings.Split(input, "\n")
	g := make(day11Grid, len(lines))
	for rowInd, line := range lines {
		g[rowInd] = make([]int, len(line))
		for colInd, char := range strings.Split(line, "") {
			var err error
			g[rowInd][colInd], err = strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
		}
	}

	return g, nil
}

type day11Grid [][]int

func (g day11Grid) evolve() int {
	for rowInd := range g {
		for colInd := range g[rowInd] {
			g[rowInd][colInd]++
		}
	}

	flashGrid := make([][]bool, len(g))
	for i := range flashGrid {
		flashGrid[i] = make([]bool, len(g[0]))
	}

	count := -1
	newCount := 0

	for count != newCount {
		count = newCount
		for rowInd := range g {
			for colInd := range g[rowInd] {
				if g[rowInd][colInd] > 9 {
					if !flashGrid[rowInd][colInd] {
						newCount++
						flashGrid[rowInd][colInd] = true
						if rowInd-1 >= 0 {
							g[rowInd-1][colInd]++
							if colInd-1 >= 0 {
								g[rowInd-1][colInd-1]++
							}
							if colInd+1 < len(g[rowInd]) {
								g[rowInd-1][colInd+1]++
							}
						}

						if rowInd+1 < len(g) {
							g[rowInd+1][colInd]++
							if colInd-1 >= 0 {
								g[rowInd+1][colInd-1]++
							}
							if colInd+1 < len(g[rowInd]) {
								g[rowInd+1][colInd+1]++
							}
						}

						if colInd-1 >= 0 {
							g[rowInd][colInd-1]++
						}

						if colInd+1 < len(g[rowInd]) {
							g[rowInd][colInd+1]++
						}
					}
				}
			}
		}
	}

	for rowInd := range g {
		for colInd := range g[rowInd] {
			if flashGrid[rowInd][colInd] {
				g[rowInd][colInd] = 0
			}
		}
	}

	return count
}

func DoDay11P1(input string) {
	sol, err := solveDay11P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 3 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay11P1(input string) (int, error) {
	grid, err := parseGrind(input)
	if err != nil {
		return 0, err
	}

	score := 0

	for i := 0; i < 100; i++ {
		score += grid.evolve()
	}

	return score, nil
}

func DoDay11P2(input string) {
	sol, err := solveDay11P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 3 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay11P2(input string) (int, error) {
	grid, err := parseGrind(input)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 1000; i++ {
		score := grid.evolve()
		if score == 100 {
			return i + 1, nil
		}
	}

	return 0, errors.New("failed to find 100 with 1000 iterations")
}
