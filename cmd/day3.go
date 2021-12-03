package cmd

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[3] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay3P1,
		P2: DoDay3P2,
	}
}

func DoDay3P1(input string) {
	sol, err := solveDay3P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 3 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay3P1(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	if len(inputs) == 0 {
		return 0, errors.New("no input")
	}

	width := len(inputs[0])
	gamma, epsilon := "", ""

	for i := 0; i < width; i++ {
		zeros, ones := d3MustCountAtLoc(inputs, i)
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaNum := d3MustParseBinaryString(gamma)
	epsilonNum := d3MustParseBinaryString(epsilon)

	return gammaNum * epsilonNum, nil
}

func DoDay3P2(input string) {
	sol, err := solveDay3P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 3 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay3P2(input string) (int, error) {
	inputs := strings.Split(input, "\n")
	if len(inputs) == 0 {
		return 0, errors.New("no input")
	}

	o2, co2 := "", ""

	o2Filter := make([]string, len(inputs))
	copy(o2Filter, inputs)
	co2Filter := make([]string, len(inputs))
	copy(co2Filter, inputs)

	for i := 0; i < len(inputs[0]); i++ {
		loc := make([]string, 0, len(o2Filter))
		zeros, ones := d3MustCountAtLoc(o2Filter, i)

		lookFor := '1'

		if zeros > ones {
			lookFor = '0'
		}

		for _, v := range o2Filter {
			if v[i] == byte(lookFor) {
				loc = append(loc, v)
			}
		}

		if len(loc) == 1 {
			o2 = loc[0]
			break
		}

		o2Filter = loc
	}

	for i := 0; i < len(inputs[0]); i++ {
		loc := make([]string, 0, len(co2Filter))
		zeros, ones := d3MustCountAtLoc(co2Filter, i)
		lookFor := '0'

		if ones < zeros {
			lookFor = '1'
		}

		for _, v := range co2Filter {
			if v[i] == byte(lookFor) {
				loc = append(loc, v)
			}
		}

		if len(loc) == 1 {
			co2 = loc[0]
			break
		}

		co2Filter = loc
	}

	if o2 == "" {
		return 0, errors.New("no o2")
	}

	if co2 == "" {
		return 0, errors.New("no co2")
	}

	o2Num := d3MustParseBinaryString(o2)

	co2Num := d3MustParseBinaryString(co2)

	return o2Num * co2Num, nil
}

func d3MustCountAtLoc(binaryStringSlice []string, loc int) (zeros, ones int) {
	for _, v := range binaryStringSlice {
		if v[loc] == '0' {
			zeros++
		} else {
			ones++
		}
	}

	return
}

func d3MustParseBinaryString(s string) int {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
