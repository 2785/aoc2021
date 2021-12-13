package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[13] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay13P1,
		P2: DoDay13P2,
	}
}

func DoDay13P1(input string) {
	sol, err := solveDay13P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 12 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

type day13Point struct{ x, y int }

func solveDay13P1(input string) (int, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return 0, errors.New("invalid input")
	}

	top, bottom := split[0], split[1]
	sparseGrid := make(map[day13Point]bool)
	for _, line := range strings.Split(top, "\n") {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}
		sparseGrid[day13Point{x, y}] = true
	}

	instructions := []struct {
		axis string
		ind  int
	}{}

	for _, line := range strings.Split(bottom, "\n") {
		line := strings.TrimPrefix(line, "fold along ")
		split := strings.Split(line, "=")
		if len(split) != 2 {
			return 0, errors.New("invalid input")
		}

		axis := split[0]
		ind, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		instructions = append(instructions, struct {
			axis string
			ind  int
		}{axis, ind})
	}

	// for _, instruction := range instructions {
	instruction := instructions[0]

	if instruction.axis == "x" {
		for k := range sparseGrid {
			if k.x > instruction.ind {
				newX := instruction.ind - (k.x - instruction.ind)
				delete(sparseGrid, k)
				sparseGrid[day13Point{newX, k.y}] = true
			}
		}
	}

	if instruction.axis == "y" {
		for k := range sparseGrid {
			if k.y > instruction.ind {
				newY := instruction.ind - (k.y - instruction.ind)
				delete(sparseGrid, k)
				sparseGrid[day13Point{k.x, newY}] = true
			}
		}
	}
	// }

	return len(sparseGrid), nil
}

func DoDay13P2(input string) {
	sol, err := solveDay13P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 12 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay13P2(input string) (int, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return 0, errors.New("invalid input")
	}

	top, bottom := split[0], split[1]
	sparseGrid := make(map[day13Point]bool)
	for _, line := range strings.Split(top, "\n") {
		split := strings.Split(line, ",")
		x, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}
		sparseGrid[day13Point{x, y}] = true
	}

	instructions := []struct {
		axis string
		ind  int
	}{}

	for _, line := range strings.Split(bottom, "\n") {
		line := strings.TrimPrefix(line, "fold along ")
		split := strings.Split(line, "=")
		if len(split) != 2 {
			return 0, errors.New("invalid input")
		}

		axis := split[0]
		ind, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}

		instructions = append(instructions, struct {
			axis string
			ind  int
		}{axis, ind})
	}

	for _, instruction := range instructions {
		if instruction.axis == "x" {
			for k := range sparseGrid {
				if k.x > instruction.ind {
					newX := instruction.ind - (k.x - instruction.ind)
					delete(sparseGrid, k)
					sparseGrid[day13Point{newX, k.y}] = true
				}
			}
		}

		if instruction.axis == "y" {
			for k := range sparseGrid {
				if k.y > instruction.ind {
					newY := instruction.ind - (k.y - instruction.ind)
					delete(sparseGrid, k)
					sparseGrid[day13Point{k.x, newY}] = true
				}
			}
		}
	}

	points := []day13Point{}
	for k := range sparseGrid {
		points = append(points, k)
	}

	xMin, yMin := points[0].x, points[0].y
	xMax, yMax := points[0].x, points[0].y
	for _, p := range points {
		if p.x < xMin {
			xMin = p.x
		}
		if p.x > xMax {
			xMax = p.x
		}
		if p.y < yMin {
			yMin = p.y
		}
		if p.y > yMax {
			yMax = p.y
		}
	}

	dx, dy := xMax-xMin, yMax-yMin
	xOffset, yOffset := xMin, yMin

	grid := make([][]string, dy+1)

	for i := 0; i < dy+1; i++ {
		grid[i] = make([]string, dx+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for k := range sparseGrid {
		grid[k.y-yOffset][k.x-xOffset] = "#"
	}

	for _, v := range grid {
		fmt.Println(strings.Join(v, ""))
	}

	return len(sparseGrid), nil
}
