package cmd

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[20] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay20P1,
		P2: DoDay20P2,
	}
}

type day20Point struct{ x, y int }

type day20Graph map[day20Point]bool

func (g day20Graph) get(p day20Point) bool {
	if thing, ok := g[p]; ok {
		return thing
	}

	return false
}

func day20EvolveGraph(g day20Graph, lookup []bool) (day20Graph, error) {
	var minX, maxX, minY, maxY int
	for k := range g {
		minX, maxX, minY, maxY = k.x, k.x, k.y, k.y
		break
	}

	for k := range g {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	minX++
	minY++
	maxX--
	maxY--

	newGraph := make(day20Graph)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			bin := ""
			if g.get(day20Point{x - 1, y - 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x, y - 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x + 1, y - 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x - 1, y}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x, y}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x + 1, y}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x - 1, y + 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x, y + 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			if g.get(day20Point{x + 1, y + 1}) {
				bin += "1"
			} else {
				bin += "0"
			}

			num, err := strconv.ParseInt(bin, 2, 64)
			if err != nil {
				return nil, err
			}

			newGraph[day20Point{x, y}] = lookup[num]
		}
	}

	return newGraph, nil
}

func DoDay20P1(input string) {
	sol, err := solveDay20P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 20 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay20P1(input string) (int, error) {
	split := strings.Split(input, "\n\n")
	lookupStr := strings.ReplaceAll(split[0], "\n", "")

	if lookupStr[0] == lookupStr[len(lookupStr)-1] {
		return 0, errors.New("wtf?")
	}

	lookup := make([]bool, len(lookupStr))
	for i, v := range lookupStr {
		if v == '#' {
			lookup[i] = true
		}
	}

	g := make(day20Graph)

	graphSplit := strings.Split(split[1], "\n")
	for y, vy := range graphSplit {
		for x, vx := range vy {
			if vx == '#' {
				g[day20Point{x, y}] = true
			} else {
				g[day20Point{x, y}] = false
			}
		}
	}

	var minX, maxX, minY, maxY int
	for k := range g {
		minX, maxX, minY, maxY = k.x, k.x, k.y, k.y
		break
	}

	for k := range g {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	minX -= 10
	maxX += 10
	minY -= 10
	maxY += 10

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if _, ok := g[day20Point{x, y}]; !ok {
				g[day20Point{x, y}] = false
			}
		}
	}

	var err error
	g, err = day20EvolveGraph(g, lookup)
	if err != nil {
		return 0, err
	}

	g, err = day20EvolveGraph(g, lookup)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, v := range g {
		if v {
			count++
		}
	}

	return count, nil
}

func DoDay20P2(input string) {
	sol, err := solveDay20P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 20 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay20P2(input string) (int, error) {
	split := strings.Split(input, "\n\n")
	lookupStr := strings.ReplaceAll(split[0], "\n", "")

	if lookupStr[0] == lookupStr[len(lookupStr)-1] {
		return 0, errors.New("wtf?")
	}

	lookup := make([]bool, len(lookupStr))
	for i, v := range lookupStr {
		if v == '#' {
			lookup[i] = true
		}
	}

	g := make(day20Graph)

	graphSplit := strings.Split(split[1], "\n")
	for y, vy := range graphSplit {
		for x, vx := range vy {
			if vx == '#' {
				g[day20Point{x, y}] = true
			} else {
				g[day20Point{x, y}] = false
			}
		}
	}

	var minX, maxX, minY, maxY int
	for k := range g {
		minX, maxX, minY, maxY = k.x, k.x, k.y, k.y
		break
	}

	for k := range g {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}

	minX -= 110
	maxX += 110
	minY -= 110
	maxY += 110

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if _, ok := g[day20Point{x, y}]; !ok {
				g[day20Point{x, y}] = false
			}
		}
	}

	for i := 0; i < 50; i++ {
		var err error

		g, err = day20EvolveGraph(g, lookup)
		if err != nil {
			return 0, err
		}

		zap.L().Info("evolving", zap.Int("iteration", i+1))
	}

	count := 0

	for _, v := range g {
		if v {
			count++
		}
	}

	return count, nil
}
