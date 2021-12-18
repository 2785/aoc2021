package cmd

import (
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[17] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay17P1,
		P2: DoDay17P2,
	}
}

func DoDay17P1(input string) {
	sol, err := solveDay17P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 17 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay17P1(input string) (int, error) {
	input = strings.TrimPrefix(input, "target area: ")
	split := strings.Split(input, ", ")
	xR := strings.TrimPrefix(split[0], "x=")
	yR := strings.TrimPrefix(split[1], "y=")
	xSplit := strings.Split(xR, "..")
	ySplit := strings.Split(yR, "..")
	xMin, err := strconv.Atoi(xSplit[0])
	if err != nil {
		return 0, err
	}
	xMax, err := strconv.Atoi(xSplit[1])
	if err != nil {
		return 0, err
	}
	yMin, err := strconv.Atoi(ySplit[0])
	if err != nil {
		return 0, err
	}
	yMax, err := strconv.Atoi(ySplit[1])
	if err != nil {
		return 0, err
	}

	maxY := 0

	for x := 1; x <= xMax; x++ {
		for y := 1; y < 10000; y++ {
			x, y := x, y
			xLoc, yLoc := 0, 0
			yMaxHeight := 0
			func() {
				for {
					xLoc, yLoc = xLoc+x, yLoc+y
					if xLoc > xMax || yLoc < yMin {
						return
					}

					if x == 0 {
						if xLoc < xMin || xLoc > xMax {
							return
						}
					}

					if yLoc <= yMaxHeight {
						if yMaxHeight <= maxY {
							return
						}
					} else {
						yMaxHeight = yLoc
					}

					if xLoc >= xMin && yLoc <= yMax {
						if yMaxHeight <= maxY {
							return
						} else {
							maxY = yMaxHeight
						}
					}

					if x > 0 {
						x--
					}

					y--
				}
			}()
		}
	}

	return maxY, nil
}

func DoDay17P2(input string) {
	sol, err := solveDay17P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 17 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay17P2(input string) (int, error) {
	input = strings.TrimPrefix(input, "target area: ")
	split := strings.Split(input, ", ")
	xR := strings.TrimPrefix(split[0], "x=")
	yR := strings.TrimPrefix(split[1], "y=")
	xSplit := strings.Split(xR, "..")
	ySplit := strings.Split(yR, "..")
	xMin, err := strconv.Atoi(xSplit[0])
	if err != nil {
		return 0, err
	}
	xMax, err := strconv.Atoi(xSplit[1])
	if err != nil {
		return 0, err
	}
	yMin, err := strconv.Atoi(ySplit[0])
	if err != nil {
		return 0, err
	}
	yMax, err := strconv.Atoi(ySplit[1])
	if err != nil {
		return 0, err
	}

	valMap := make(map[struct{ x, y int }]bool)

	for xLoop := 1; xLoop <= xMax; xLoop++ {
		for yLoop := yMin; yLoop < 10000; yLoop++ {
			x, y := xLoop, yLoop
			xLoc, yLoc := 0, 0
			func() {
				for {
					xLoc, yLoc = xLoc+x, yLoc+y
					if xLoc > xMax || yLoc < yMin {
						return
					}

					if x == 0 {
						if xLoc < xMin || xLoc > xMax {
							return
						}
					}

					if xLoc >= xMin && yLoc <= yMax {
						valMap[struct {
							x int
							y int
						}{xLoop, yLoop}] = true
					}

					if x > 0 {
						x--
					}

					y--
				}
			}()
		}
	}

	return len(valMap), nil
}
