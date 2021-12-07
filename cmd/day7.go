package cmd

import (
	"strings"

	"github.com/2785/aoc2021/util"
	"go.uber.org/zap"
)

func init() {
	solvers[7] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay7P1,
		P2: DoDay7P2,
	}
}

func DoDay7P1(input string) {
	sol, err := solveDay7P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 7 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay7P1(input string) (int, error) {
	inputs := strings.Split(input, ",")
	nums, err := util.IntsFromStrings(inputs)
	if err != nil {
		return 0, err
	}

	min, max := 0, 0
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	minFuel := 0

	for _, v := range nums {
		dist := v
		if dist < 0 {
			dist = -dist
		}

		minFuel += dist
	}

	for i := min + 1; i <= max; i++ {
		fuel := 0
		for _, v := range nums {
			dist := v - i
			if dist < 0 {
				dist = -dist
			}

			fuel += dist
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel, nil
}

func DoDay7P2(input string) {
	sol, err := solveDay7P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 7 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay7P2(input string) (int, error) {
	inputs := strings.Split(input, ",")
	nums, err := util.IntsFromStrings(inputs)
	if err != nil {
		return 0, err
	}

	min, max := 0, 0
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	minFuel := 0

	for _, v := range nums {
		dist := v
		if dist < 0 {
			dist = -dist
		}

		minFuel += dist * (dist + 1) / 2
	}

	for i := min + 1; i <= max; i++ {
		fuel := 0
		for _, v := range nums {
			dist := v - i
			if dist < 0 {
				dist = -dist
			}

			fuel += dist * (dist + 1) / 2
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel, nil
}
