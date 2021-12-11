package cmd

import (
	"errors"
	"sort"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[10] = struct {
		P1 func(string)
		P2 func(string)
	}{
		DoDay10P1, DoDay10P2,
	}
}

var day10Mapping = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var day10ReverseMapping = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var day10ScoreMapping = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var day10AutoCompleteScoreMapping = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func SolveDay10P1(input string) (int, error) {
	lines := strings.Split(input, "\n")

	score := 0

Line:
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		s := make(day10Stack, 0)
		for _, char := range splitLine {
			if inKeys(day10Mapping, char) {
				s.Push(char)
			} else {
				if err := s.Pop(char); err != nil {
					score += day10ScoreMapping[char]
					continue Line
				}
			}
		}
	}

	return score, nil
}

func DoDay10P1(input string) {
	sol, err := SolveDay10P1(input)
	if err != nil {
		zap.L().Fatal("failed to solve day 10 p1", zap.Error(err))
		return
	}

	zap.L().Info("result", zap.Int("count", sol), zap.Int("part", 1))
}

func SolveDay10P2(input string) (int, error) {
	lines := strings.Split(input, "\n")

	welFormed := make([]day10Stack, 0)

Line:
	for _, line := range lines {
		splitLine := strings.Split(line, "")
		s := make(day10Stack, 0)
		for _, char := range splitLine {
			if inKeys(day10Mapping, char) {
				s.Push(char)
			} else {
				if err := s.Pop(char); err != nil {
					continue Line
				}
			}
		}

		welFormed = append(welFormed, s)
	}

	scores := make([]int, 0)
	for _, s := range welFormed {
		score := 0
		for i := len(s) - 1; i >= 0; i-- {
			score *= 5
			score += day10AutoCompleteScoreMapping[day10Mapping[s[i]]]
		}

		scores = append(scores, score)
	}

	if len(scores)%2 == 0 {
		return 0, errors.New("wtf?")
	}

	sort.Ints(scores)

	return scores[len(scores)/2], nil

}

func DoDay10P2(input string) {
	sol, err := SolveDay10P2(input)
	if err != nil {
		zap.L().Fatal("failed to solve day 10 p2", zap.Error(err))
		return
	}

	zap.L().Info("result", zap.Int("count", sol), zap.Int("part", 2))
}

type day10Stack []string

func (s *day10Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *day10Stack) Pop(v string) error {
	if (*s)[len(*s)-1] != day10ReverseMapping[v] {
		return errors.New("invalid stack")
	}

	*s = (*s)[:len(*s)-1]
	return nil
}

func inKeys(mapping map[string]string, key string) bool {
	for k := range mapping {
		if k == key {
			return true
		}
	}
	return false
}
