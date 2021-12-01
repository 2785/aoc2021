package cmd

import (
	"strings"

	"github.com/2785/aoc2021/util"
	"go.uber.org/zap"
)

func init() {
	solvers[1] = struct {
		P1 func(string)
		P2 func(string)
	}{
		DoDay1P1, DoDay1P2,
	}
}

func SolveDay1P1(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	count := 0
	nums, err := util.IntsFromStrings(inputs)
	if err != nil {
		return 0, err
	}

	for i, num := range nums {
		if i == 0 {
			continue
		}

		if num > nums[i-1] {
			count++
		}
	}

	return count, nil
}

func DoDay1P1(input string) {
	sol, err := SolveDay1P1(input)
	if err != nil {
		zap.L().Error("failed to solve day 1 p1", zap.Error(err))
		return
	}

	zap.L().Info("result", zap.Int("count", sol), zap.Int("part", 1))
}

func SolveDay1P2(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	count := 0

	nums, err := util.IntsFromStrings(inputs)
	if err != nil {
		return 0, err
	}

	sums := make([]int, len(nums)-2)

	for i := 0; i < len(nums)-2; i++ {
		sums[i] = nums[i] + nums[i+1] + nums[i+2]
	}

	for i, sum := range sums {
		if i == 0 {
			continue
		}

		if sum > sums[i-1] {
			count++
		}
	}

	return count, nil
}

func DoDay1P2(input string) {
	sol, err := SolveDay1P2(input)
	if err != nil {
		zap.L().Error("failed to solve day 1 p2", zap.Error(err))
		return
	}

	zap.L().Info("result", zap.Int("count", sol), zap.Int("part", 2))
}
