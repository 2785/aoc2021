package cmd

import (
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[22] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay22P1,
		P2: DoDay22P2,
	}
}

func DoDay22P1(input string) {
	sol, err := solveDay22P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 22 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

type day22Point struct {
	x, y, z int
}

type day22Instruction struct {
	minx, maxx, miny, maxy, minz, maxz int
	state                              bool
}

type day22Block struct {
	minx, maxx, miny, maxy, minz, maxz int
}

func deductBlock(src day22Block, thing day22Block) []day22Block {
	if src.minx > thing.maxx || src.maxx < thing.minx || src.miny > thing.maxy || src.maxy < thing.miny || src.minz > thing.maxz || src.maxz < thing.minz {
		return []day22Block{src}
	}

	xmin, xmax, ymin, ymax, zmin, zmax := src.minx, src.maxx, src.miny, src.maxy, src.minz, src.maxz

	if thing.minx > xmin {
		xmin = thing.minx
	}

	if thing.maxx < xmax {
		xmax = thing.maxx
	}

	if thing.miny > ymin {
		ymin = thing.miny
	}

	if thing.maxy < ymax {
		ymax = thing.maxy
	}

	if thing.minz > zmin {
		zmin = thing.minz
	}

	if thing.maxz < zmax {
		zmax = thing.maxz
	}

	blocks := []day22Block{}
	// need to deal with all the blocks that potentially surround the intersect block

	if xmin > src.minx {
		blocks = append(blocks, day22Block{src.minx, xmin - 1, src.miny, src.maxy, src.minz, src.maxz})
	}

	if xmax < src.maxx {
		blocks = append(blocks, day22Block{xmax + 1, src.maxx, src.miny, src.maxy, src.minz, src.maxz})
	}

	if ymin > src.miny {
		blocks = append(blocks, day22Block{xmin, xmax, src.miny, ymin - 1, src.minz, src.maxz})
	}

	if ymax < src.maxy {
		blocks = append(blocks, day22Block{xmin, xmax, ymax + 1, src.maxy, src.minz, src.maxz})
	}

	if zmin > src.minz {
		blocks = append(blocks, day22Block{xmin, xmax, ymin, ymax, src.minz, zmin - 1})
	}

	if zmax < src.maxz {
		blocks = append(blocks, day22Block{xmin, xmax, ymin, ymax, zmax + 1, src.maxz})
	}

	return blocks
}

func day22ParseLine(line string) (day22Instruction, error) {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		return day22Instruction{}, errors.New("expected 2 sections")
	}
	inst := day22Instruction{}
	if split[0] == "on" {
		inst.state = true
	}

	line = split[1]
	split = strings.Split(line, ",")
	if len(split) != 3 {
		return day22Instruction{}, errors.New("expected 3 coordinates")
	}

	xSplit := strings.Split(split[0], "=")
	xSplit = strings.Split(xSplit[1], "..")
	minx, err := strconv.Atoi(xSplit[0])
	if err != nil {
		return day22Instruction{}, err
	}

	maxx, err := strconv.Atoi(xSplit[1])
	if err != nil {
		return day22Instruction{}, err
	}

	ySplit := strings.Split(split[1], "=")
	ySplit = strings.Split(ySplit[1], "..")
	miny, err := strconv.Atoi(ySplit[0])
	if err != nil {
		return day22Instruction{}, err
	}

	maxy, err := strconv.Atoi(ySplit[1])
	if err != nil {
		return day22Instruction{}, err
	}

	zSplit := strings.Split(split[2], "=")
	zSplit = strings.Split(zSplit[1], "..")
	minz, err := strconv.Atoi(zSplit[0])
	if err != nil {
		return day22Instruction{}, err
	}

	maxz, err := strconv.Atoi(zSplit[1])
	if err != nil {
		return day22Instruction{}, err
	}

	inst.minx = minx
	inst.maxx = maxx
	inst.miny = miny
	inst.maxy = maxy
	inst.minz = minz
	inst.maxz = maxz

	return inst, nil
}

func solveDay22P1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	instructions := make([]day22Instruction, len(lines))
	for i, line := range lines {
		inst, err := day22ParseLine(line)
		if err != nil {
			return 0, err
		}
		instructions[i] = inst
	}

	min, max := -50, 50
	grid := make(map[day22Point]bool)

	for _, inst := range instructions {
		inst := inst
		if inst.minx < min {
			inst.minx = min
		}
		if inst.maxx > max {
			inst.maxx = max
		}
		if inst.miny < min {
			inst.miny = min
		}
		if inst.maxy > max {
			inst.maxy = max
		}
		if inst.minz < min {
			inst.minz = min
		}
		if inst.maxz > max {
			inst.maxz = max
		}

		for x := inst.minx; x <= inst.maxx; x++ {
			for y := inst.miny; y <= inst.maxy; y++ {
				for z := inst.minz; z <= inst.maxz; z++ {
					if x < min || x > max || y < min || y > max || z < min || z > max {
						continue
					}

					if inst.state {
						grid[day22Point{x, y, z}] = true
					} else {
						delete(grid, day22Point{x, y, z})
					}
				}
			}
		}
	}

	return len(grid), nil

}

func DoDay22P2(input string) {
	sol, err := solveDay22P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 22 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay22P2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	instructions := make([]day22Instruction, len(lines))
	for i, line := range lines {
		inst, err := day22ParseLine(line)
		if err != nil {
			return 0, err
		}
		instructions[i] = inst
	}

	totalOn := 0

	for i := len(instructions) - 1; i >= 0; i-- {
		inst := instructions[i]
		if !inst.state {
			continue
		}

		blocks := []day22Block{{
			minx: inst.minx,
			maxx: inst.maxx,
			miny: inst.miny,
			maxy: inst.maxy,
			minz: inst.minz,
			maxz: inst.maxz,
		}}

		for j := i + 1; j < len(instructions); j++ {
			target := instructions[j]
			blockToRemove := day22Block{
				minx: target.minx,
				maxx: target.maxx,
				miny: target.miny,
				maxy: target.maxy,
				minz: target.minz,
				maxz: target.maxz,
			}

			newBlocks := make([]day22Block, 0)

			for _, block := range blocks {
				newBlocks = append(newBlocks, deductBlock(block, blockToRemove)...)
			}

			blocks = newBlocks
		}

		for _, block := range blocks {
			totalOn += (block.maxx - block.minx + 1) * (block.maxy - block.miny + 1) * (block.maxz - block.minz + 1)
		}
	}

	return totalOn, nil
}
