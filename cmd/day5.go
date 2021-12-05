package cmd

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func init() {
	solvers[5] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay5P1,
		P2: DoDay5P2,
	}
}

func DoDay5P1(input string) {
	sol, err := solveDay5P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 5 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay5P1(input string) (int, error) {
	lines, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	maxX, maxY := getGridSize(lines)

	grid := make([][]int, maxY)
	for i := range grid {
		grid[i] = make([]int, maxX)
	}

	for _, line := range lines {
		if line.A.X == line.B.X {
			yMin, yMax := line.A.Y, line.B.Y
			if yMin > yMax {
				yMin, yMax = yMax, yMin
			}

			for y := yMin; y <= yMax; y++ {
				grid[y][line.A.X] += 1
			}
		}

		if line.A.Y == line.B.Y {
			xMin, xMax := line.A.X, line.B.X
			if xMin > xMax {
				xMin, xMax = xMax, xMin
			}

			for x := xMin; x <= xMax; x++ {
				grid[line.A.Y][x] += 1
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	return count, nil
}

func DoDay5P2(input string) {
	sol, err := solveDay5P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 5 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay5P2(input string) (int, error) {
	lines, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	maxX, maxY := getGridSize(lines)

	grid := make([][]int, maxY)
	for i := range grid {
		grid[i] = make([]int, maxX)
	}

	for _, line := range lines {
		if line.A.X == line.B.X {
			yMin, yMax := line.A.Y, line.B.Y
			if yMin > yMax {
				yMin, yMax = yMax, yMin
			}

			for y := yMin; y <= yMax; y++ {
				grid[y][line.A.X] += 1
			}

			continue
		}

		if line.A.Y == line.B.Y {
			xMin, xMax := line.A.X, line.B.X
			if xMin > xMax {
				xMin, xMax = xMax, xMin
			}

			for x := xMin; x <= xMax; x++ {
				grid[line.A.Y][x] += 1
			}

			continue
		}

		dx, dy := line.B.X-line.A.X, line.B.Y-line.A.Y
		if dx < 0 {
			dx = -dx
		}

		if dy < 0 {
			dy = -dy
		}

		if dx == dy {
			xStart := line.A.X
			yStart := line.A.Y

			xSpan := 1
			if xStart > line.B.X {
				xSpan = -1
			}

			ySpan := 1
			if yStart > line.B.Y {
				ySpan = -1
			}

			for d := 0; d <= dx; d++ {
				grid[yStart+d*ySpan][xStart+d*xSpan] += 1
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	return count, nil
}

type d5Coord struct {
	X, Y int
}

type d5Line struct {
	A, B d5Coord
}

func parseInput(input string) ([]d5Line, error) {
	lines := strings.Split(input, "\n")

	var result []d5Line
	for _, line := range lines {
		split := strings.Split(line, "->")
		if len(split) != 2 {
			return nil, errors.New("invalid input")
		}

		a := strings.Split(strings.TrimSpace(split[0]), ",")
		if len(a) != 2 {
			return nil, errors.New("invalid input")
		}

		b := strings.Split(strings.TrimSpace(split[1]), ",")
		if len(b) != 2 {
			return nil, errors.New("invalid input")
		}

		x1, err := strconv.Atoi(a[0])
		if err != nil {
			return nil, err
		}

		y1, err := strconv.Atoi(a[1])
		if err != nil {
			return nil, err
		}

		x2, err := strconv.Atoi(b[0])
		if err != nil {
			return nil, err
		}

		y2, err := strconv.Atoi(b[1])
		if err != nil {
			return nil, err
		}

		result = append(result, d5Line{
			A: d5Coord{
				X: x1,
				Y: y1,
			},
			B: d5Coord{
				X: x2,
				Y: y2,
			},
		})
	}

	return result, nil
}

func getGridSize(grid []d5Line) (x int, y int) {
	maxX, maxY := 0, 0

	for _, line := range grid {
		if line.A.X > maxX {
			maxX = line.A.X
		}

		if line.A.Y > maxY {
			maxY = line.A.Y
		}

		if line.B.X > maxX {
			maxX = line.B.X
		}

		if line.B.Y > maxY {
			maxY = line.B.Y
		}
	}

	return maxX + 1, maxY + 1
}
