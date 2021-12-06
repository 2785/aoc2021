package cmd

import (
	"strings"

	"github.com/2785/aoc2021/util"
	"go.uber.org/zap"
)

func init() {
	solvers[6] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay6P1,
		P2: DoDay6P2,
	}
}

func DoDay6P1(input string) {
	sol, err := solveDay6(input, 80)
	if err != nil {
		zap.L().Fatal("Failed to solve day 6 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func DoDay6P2(input string) {
	sol, err := solveDay6(input, 256)
	if err != nil {
		zap.L().Fatal("Failed to solve day 6 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay6(input string, days int) (int, error) {
	inputs := strings.Split(input, ",")
	fishes, err := util.IntsFromStrings(inputs)
	if err != nil {
		return 0, err
	}

	sequence := make([]int, 9)
	for _, v := range fishes {
		sequence[v]++
	}

	for i := 0; i < days; i++ {
		reproduce := sequence[0]
		for j := 1; j < len(sequence); j++ {
			sequence[j-1] = sequence[j]
		}
		sequence[6] += reproduce
		sequence[8] = reproduce
	}

	sum := 0
	for _, v := range sequence {
		sum += v
	}

	return sum, nil
}
