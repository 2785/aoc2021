package cmd

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func init() {
	solvers[4] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay4P1,
		P2: DoDay4P2,
	}
}

func DoDay4P1(input string) {
	sol, err := solveDay4P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 4 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay4P1(input string) (int, error) {
	segments := strings.Split(input, "\n\n")

	if len(segments) < 2 {
		return 0, errors.New("invalid input")
	}

	pool := strings.Split(segments[0], ",")

	boards := make([]d4Board, len(segments)-1)

	for i := 1; i < len(segments); i++ {
		b, err := parseBoard(segments[i])
		if err != nil {
			return 0, errors.Wrap(err, "failed to parse board")
		}
		boards[i-1] = b
	}

	for _, v := range pool {
		for _, b := range boards {
			b.mark(v)

			if b.wins() {
				return b.score(v)
			}
		}
	}

	return 0, errors.New("no solution found")
}

func DoDay4P2(input string) {
	sol, err := solveDay4P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 4 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay4P2(input string) (int, error) {
	segments := strings.Split(input, "\n\n")

	if len(segments) < 2 {
		return 0, errors.New("invalid input")
	}

	pool := strings.Split(segments[0], ",")

	boards := make([]d4Board, len(segments)-1)

	for i := 1; i < len(segments); i++ {
		b, err := parseBoard(segments[i])
		if err != nil {
			return 0, errors.Wrap(err, "failed to parse board")
		}
		boards[i-1] = b
	}

	boardCount := len(boards)

	for _, v := range pool {
		newBoard := make([]d4Board, 0, len(segments)-1)

		for _, b := range boards {
			b.mark(v)

			if b.wins() {
				if boardCount == 1 {
					return b.score(v)
				} else {
					boardCount--
				}
			} else {
				newBoard = append(newBoard, b)
			}
		}

		boards = newBoard
	}

	return 0, errors.New("no solution found")
}

type d4Entry struct {
	num    string
	marked bool
}

type d4Board [][]d4Entry

func (b d4Board) valid() error {
	if len(b) != len(b[0]) {
		return errors.New("invalid board")
	}

	for _, v := range b {
		if len(v) != len(b[0]) {
			return errors.New("invalid board")
		}
	}

	return nil
}

func (b d4Board) wins() bool {
	for _, v := range b {
		rowAllGood := true
		for _, v2 := range v {
			if !v2.marked {
				rowAllGood = false
				break
			}
		}

		if rowAllGood {
			return true
		}
	}

	for i := range b[0] {
		colAllGood := true
		for j := range b {
			if !b[j][i].marked {
				colAllGood = false
				break
			}
		}

		if colAllGood {
			return true
		}
	}

	return false
}

func (b d4Board) score(target string) (int, error) {
	targetNum, err := strconv.Atoi(target)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse number")
	}

	score := 0
	for _, v := range b {
		for _, v2 := range v {
			if !v2.marked {
				num, err := strconv.Atoi(v2.num)
				if err != nil {
					return 0, errors.Wrap(err, "failed to parse number")
				}
				score += num
			}
		}
	}

	return score * targetNum, nil
}

func (b d4Board) mark(target string) {
	for i := range b {
		for j := range b[i] {
			if b[i][j].num == target {
				b[i][j].marked = true
			}
		}
	}
}

func parseBoard(boardInput string) (d4Board, error) {
	lines := strings.Split(boardInput, "\n")
	board := make(d4Board, len(lines))
	for j, v := range lines {
		board[j] = parseLine(v, len(lines))
	}

	if err := board.valid(); err != nil {
		return nil, err
	}

	return board, nil
}

func parseLine(line string, count int) []d4Entry {
	entries := make([]d4Entry, count)
	for i := 0; i < count; i++ {
		entries[i].num = strings.TrimSpace(line[i*3 : i*3+2])
	}

	return entries
}
