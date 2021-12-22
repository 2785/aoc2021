package cmd

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[21] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay21P1,
		P2: DoDay21P2,
	}
}

func DoDay21P1(input string) {
	sol, err := solveDay21P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 21 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay21P1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	if len(lines) != 2 {
		return 0, errors.New("expected 2 lines")
	}

	p1, err := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {
		return 0, err
	}

	p2, err := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {
		return 0, err
	}

	p1--
	p2--

	p1Score, p2Score := 0, 0

	dieTracker := 1
	rollCount := 0
	roll := func() int {
		num := dieTracker
		dieTracker++
		if dieTracker > 100 {
			dieTracker = 1
		}

		rollCount++
		return num
	}

	for {
		p1 += roll() + roll() + roll()
		p1 %= 10
		p1Score += p1 + 1

		if p1Score >= 1000 {
			break
		}

		p2 += roll() + roll() + roll()
		p2 %= 10
		p2Score += p2 + 1

		if p2 >= 1000 {
			break
		}
	}

	loser := p1Score
	if p2Score < p1Score {
		loser = p2Score
	}

	return loser * rollCount, nil
}

func DoDay21P2(input string) {
	sol, err := solveDay21P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 21 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay21P2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	if len(lines) != 2 {
		return 0, errors.New("expected 2 lines")
	}

	p1, err := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {
		return 0, err
	}

	p2, err := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {
		return 0, err
	}

	p1--
	p2--

	worldLookup := map[int]int{}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				sum := i + j + k
				if _, ok := worldLookup[sum]; !ok {
					worldLookup[sum] = 0
				}

				worldLookup[sum]++
			}
		}
	}

	w1, w2 := 0, 0

	type day21Node struct {
		worldCount       int
		p1Score, p2Score int
		p1Loc, p2Loc     int
		roll             int
		p1NextTurn       bool
	}

	nodeWithRoll := func(node day21Node, roll int) day21Node {
		node.roll = roll
		node.worldCount = worldLookup[roll] * node.worldCount
		if node.p1NextTurn {
			node.p1Loc = (node.p1Loc + roll) % 10
			node.p1Score += node.p1Loc + 1
			node.p1NextTurn = false
		} else {
			node.p2Loc = (node.p2Loc + roll) % 10
			node.p2Score += node.p2Loc + 1
			node.p1NextTurn = true
		}

		return node
	}

	nodeWins := func(node day21Node) bool {
		if node.p1Score >= 21 {
			w1 += node.worldCount
			return true
		} else if node.p2Score >= 21 {
			w2 += node.worldCount
			return true
		}

		return false
	}

	stack := []day21Node{{
		worldCount: 1,
		p1Score:    0,
		p2Score:    0,
		p1Loc:      p1,
		p2Loc:      p2,
		p1NextTurn: true,
	}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for k := range worldLookup {
			newNode := nodeWithRoll(node, k)
			if !nodeWins(newNode) {
				stack = append(stack, newNode)
			}
		}
	}

	if w1 > w2 {
		return w1, nil
	} else {
		return w2, nil
	}
}
