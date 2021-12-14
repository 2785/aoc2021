package cmd

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[14] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay14P1,
		P2: DoDay14P2,
	}
}

func DoDay14P1(input string) {
	sol, err := solveDay14P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 14 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay14P1(input string) (int, error) {
	parts := strings.Split(input, "\n\n")

	if len(parts) != 2 {
		return 0, errors.New("invalid input")
	}

	top, bottom := parts[0], parts[1]

	thing := strings.Split(top, "")

	ruleSplit := strings.Split(bottom, "\n")
	rules := make(map[string]string, len(ruleSplit))

	for _, line := range ruleSplit {
		split := strings.Split(line, " -> ")
		if len(split) != 2 {
			return 0, errors.New("invalid rule")
		}
		rules[split[0]] = split[1]
	}

	for i := 0; i < 10; i++ {
		insertions := []struct {
			s   string
			ind int
		}{}

		offSet := 0

		for i := 0; i < len(thing)-1; i++ {
			seg := thing[i : i+2]
			if rule, ok := rules[strings.Join(seg, "")]; ok {
				insertions = append(insertions, struct {
					s   string
					ind int
				}{s: rule, ind: i + 1})
			}
		}

		for _, ins := range insertions {
			thing = append(thing[:ins.ind+offSet], append([]string{ins.s}, thing[ins.ind+offSet:]...)...)
			offSet++
		}
	}

	counts := make(map[string]int)
	for _, v := range thing {
		if _, ok := counts[v]; !ok {
			counts[v] = 0
		}

		counts[v]++
	}

	max, min := 0, len(thing)
	for _, v := range counts {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max - min, nil
}

func DoDay14P2(input string) {
	sol, err := solveDay14P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 14 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay14P2(input string) (int, error) {
	parts := strings.Split(input, "\n\n")

	if len(parts) != 2 {
		return 0, errors.New("invalid input")
	}

	top, bottom := parts[0], parts[1]

	thing := strings.Split(top, "")

	ruleSplit := strings.Split(bottom, "\n")
	rules := make(map[string]string, len(ruleSplit))

	for _, line := range ruleSplit {
		split := strings.Split(line, " -> ")
		if len(split) != 2 {
			return 0, errors.New("invalid rule")
		}
		rules[split[0]] = split[1]
	}

	pairMap := make(map[string]int)

	for i := 0; i < len(thing)-1; i++ {
		pair := strings.Join(thing[i:i+2], "")

		if _, ok := pairMap[pair]; !ok {
			pairMap[pair] = 0
		}

		pairMap[pair]++
	}

	letterCount := make(map[string]int)

	for _, v := range thing {
		if _, ok := letterCount[v]; !ok {
			letterCount[v] = 0
		}

		letterCount[v]++
	}

	for i := 0; i < 40; i++ {
		newMap := make(map[string]int)
		for k, v := range pairMap {
			newMap[k] = v
		}

		for pair, count := range pairMap {
			if rule, ok := rules[pair]; ok {
				newMap[pair] -= count
				left, right := pair[:1], pair[1:]
				left += rule
				right = rule + right
				if _, ok := pairMap[left]; !ok {
					pairMap[left] = 0
				}
				newMap[left] += count
				if _, ok := newMap[right]; !ok {
					pairMap[right] = 0
				}
				newMap[right] += count
				if _, ok := letterCount[rule]; !ok {
					letterCount[rule] = 0
				}
				letterCount[rule] += count
			}
		}

		pairMap = newMap
	}

	max, min := 0, letterCount[thing[0]]
	for _, v := range letterCount {
		if v > max {
			max = v
		}

		if v != 0 && v < min {
			min = v
		}
	}

	return max - min, nil
}
