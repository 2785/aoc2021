package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

const (
	fwd = "forward"
	dwn = "down"
	up  = "up"
)

func init() {
	solvers[2] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay2P1,
		P2: DoDay2P2,
	}
}

func DoDay2P1(input string) {
	sol, err := solveDay2P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 2 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay2P1(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	depth, dist := 0, 0
	for _, i := range inputs {
		split := strings.Split(i, " ")
		if len(split) != 2 {
			return 0, fmt.Errorf("invalid input: %s", i)
		}

		numStr := split[1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		switch split[0] {
		case fwd:
			dist += num
		case dwn:
			depth += num
		case up:
			depth -= num
		default:
			return 0, fmt.Errorf("invalid direction: %s", split[0])
		}
	}

	return depth * dist, nil
}

func DoDay2P2(input string) {
	sol, err := solveDay2P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 2 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay2P2(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	depth, dist, aim := 0, 0, 0

	for _, i := range inputs {
		split := strings.Split(i, " ")
		if len(split) != 2 {
			return 0, fmt.Errorf("invalid input: %s", i)
		}

		numStr := split[1]

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		switch split[0] {
		case fwd:
			dist += num
			depth += num * aim
		case dwn:
			aim += num
		case up:
			aim -= num
		default:
			return 0, fmt.Errorf("invalid direction: %s", split[0])
		}
	}

	return depth * dist, nil
}
