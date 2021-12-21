package cmd

import (
	"sort"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[19] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay19P1,
		P2: DoDay19P2,
	}
}

type day19Point struct {
	x, y, z int
}

type day19Points []day19Point

func (p day19Points) Len() int {
	return len(p)
}

func (p day19Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p day19Points) Less(i, j int) bool {
	if p[i].x != p[j].x {
		return p[i].x < p[j].x
	}

	if p[i].y != p[j].y {
		return p[i].y < p[j].y
	}

	return p[i].z < p[j].z
}

var day19ConfigurationMappers = []func(day19Point) day19Point{
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.x,
			y: dp.y,
			z: dp.z,
		}
	},
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.x,
			y: dp.z,
			z: dp.y,
		}
	},
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.y,
			y: dp.x,
			z: dp.z,
		}
	},
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.y,
			y: dp.z,
			z: dp.x,
		}
	},
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.z,
			y: dp.y,
			z: dp.x,
		}
	},
	func(dp day19Point) day19Point {
		return day19Point{
			x: dp.z,
			y: dp.x,
			z: dp.y,
		}
	},
}

var day19DirectionMappers = []func(day19Point) day19Point{}

func init() {
	for _, x := range []int{-1, 1} {
		for _, y := range []int{-1, 1} {
			for _, z := range []int{-1, 1} {
				x, y, z := x, y, z
				day19DirectionMappers = append(day19DirectionMappers, func(dp day19Point) day19Point {
					return day19Point{
						x: dp.x * x,
						y: dp.y * y,
						z: dp.z * z,
					}
				})
			}
		}
	}
}

type day19Scanner struct {
	rawPoints day19Points
}

func (s *day19Scanner) getConfigurations() []day19Points {
	configBases := make([][]day19Point, len(day19ConfigurationMappers))
	for i, mapper := range day19ConfigurationMappers {
		configBases[i] = make([]day19Point, len(s.rawPoints))
		for j, dp := range s.rawPoints {
			configBases[i][j] = mapper(dp)
		}
	}

	allConfigs := make([]day19Points, 0, len(configBases)*len(day19DirectionMappers))
	for _, configBase := range configBases {
		for _, directionMapper := range day19DirectionMappers {
			config := make([]day19Point, len(configBase))
			for i, dp := range configBase {
				config[i] = directionMapper(dp)
			}
			allConfigs = append(allConfigs, config)
		}
	}

	return allConfigs
}

func day19ParseScanners(input string) ([]day19Scanner, error) {
	scannerStrings := strings.Split(input, "\n\n")
	scanners := make([]day19Scanner, len(scannerStrings))
	for i, scannerString := range scannerStrings {
		scanner, err := day19ParseScanner(scannerString)
		if err != nil {
			return nil, err
		}

		scanners[i] = scanner
	}

	return scanners, nil
}

func day19ParseScanner(input string) (day19Scanner, error) {
	lines := strings.Split(input, "\n")
	lines = lines[1:]

	points := make(day19Points, len(lines))
	for i, line := range lines {
		point, err := day19ParsePoint(line)
		if err != nil {
			return day19Scanner{}, err
		}

		points[i] = point
	}

	return day19Scanner{
		rawPoints: points,
	}, nil
}

func day19ParsePoint(line string) (day19Point, error) {
	pointSplit := strings.Split(line, ",")
	x, err := strconv.Atoi(pointSplit[0])
	if err != nil {
		return day19Point{}, err
	}
	y, err := strconv.Atoi(pointSplit[1])
	if err != nil {
		return day19Point{}, err
	}
	z, err := strconv.Atoi(pointSplit[2])
	if err != nil {
		return day19Point{}, err
	}

	return day19Point{
		x: x,
		y: y,
		z: z,
	}, nil
}

func day19CompareConfiguration(config1, config2 day19Points, threshold int) (match bool, offset day19Point) {
	config1Loc := make(day19Points, len(config1))
	copy(config1Loc, config1)
	sort.Sort(config1Loc)

	config2Loc := make(day19Points, len(config2))
	copy(config2Loc, config2)
	sort.Sort(config2Loc)

	for _, lhs := range config1Loc {
		for _, rhs := range config2Loc {
			offset = day19Point{
				x: rhs.x - lhs.x,
				y: rhs.y - lhs.y,
				z: rhs.z - lhs.z,
			}

			dupCounter := 0
			for _, left := range config1Loc {
				for _, right := range config2Loc {
					offsetted := day19Point{
						x: left.x + offset.x,
						y: left.y + offset.y,
						z: left.z + offset.z,
					}

					if offsetted.x > 1000 || offsetted.x < -1000 || offsetted.y > 1000 || offsetted.y < -1000 || offsetted.z > 1000 || offsetted.z < -1000 {
						// out of range
						continue
					}

					if offsetted == right {
						dupCounter++
					}
				}
			}

			if dupCounter >= threshold {
				return true, day19Point{
					x: -offset.x,
					y: -offset.y,
					z: -offset.z,
				}
			}
		}
	}

	return false, offset
}

func day19CompareScanners(scanner1, scanner2 day19Scanner, threshold int) (match bool, offset day19Point, configuration day19Points) {
	s2Configs := scanner2.getConfigurations()

	for _, config2 := range s2Configs {
		match, offset = day19CompareConfiguration(scanner1.rawPoints, config2, threshold)
		if match {
			return match, offset, config2
		}
	}

	return false, offset, configuration
}

func DoDay19P1(input string) {
	sol, err := solveDay19P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 19 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay19P1(input string) (int, error) {
	scanners, err := day19ParseScanners(input)
	if err != nil {
		return 0, err
	}

	normalizedScanners := map[day19Point]day19Scanner{
		{x: 0, y: 0, z: 0}: scanners[0],
	}
	scanners = scanners[1:]

	for {
		remainingScanners := make([]day19Scanner, 0, len(scanners))
		for _, scanner := range scanners {
			scanner := scanner
			matched := false
			for leftOffset, normalizedScanner := range normalizedScanners {
				match, offset, config := day19CompareScanners(normalizedScanner, scanner, 12)
				if match {
					matched = true
					offset = day19Point{
						x: leftOffset.x + offset.x,
						y: leftOffset.y + offset.y,
						z: leftOffset.z + offset.z,
					}

					normalizedScanners[offset] = day19Scanner{
						rawPoints: config,
					}

					break
				}
			}

			if !matched {
				remainingScanners = append(remainingScanners, scanner)
			}
		}

		if len(remainingScanners) == 0 {
			break
		}

		scanners = remainingScanners
	}

	beacons := make(map[day19Point]bool)

	for offset, scanner := range normalizedScanners {
		for _, point := range scanner.rawPoints {
			beacon := day19Point{
				x: point.x + offset.x,
				y: point.y + offset.y,
				z: point.z + offset.z,
			}

			beacons[beacon] = true
		}
	}

	return len(beacons), nil
}

func DoDay19P2(input string) {
	sol, err := solveDay19P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 19 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay19P2(input string) (int, error) {
	scanners, err := day19ParseScanners(input)
	if err != nil {
		return 0, err
	}

	normalizedScanners := map[day19Point]day19Scanner{
		{x: 0, y: 0, z: 0}: scanners[0],
	}
	scanners = scanners[1:]

	for {
		remainingScanners := make([]day19Scanner, 0, len(scanners))
		for _, scanner := range scanners {
			scanner := scanner
			matched := false
			for leftOffset, normalizedScanner := range normalizedScanners {
				match, offset, config := day19CompareScanners(normalizedScanner, scanner, 12)
				if match {
					matched = true
					offset = day19Point{
						x: leftOffset.x + offset.x,
						y: leftOffset.y + offset.y,
						z: leftOffset.z + offset.z,
					}

					normalizedScanners[offset] = day19Scanner{
						rawPoints: config,
					}

					break
				}
			}

			if !matched {
				remainingScanners = append(remainingScanners, scanner)
			}
		}

		if len(remainingScanners) == 0 {
			break
		}

		scanners = remainingScanners
	}

	offsetSlice := make([]day19Point, 0, len(normalizedScanners))
	for k := range normalizedScanners {
		offsetSlice = append(offsetSlice, k)
	}

	maxDist := 0

	for i := 0; i < len(offsetSlice); i++ {
		for j := i + 1; j < len(offsetSlice); j++ {
			dx, dy, dz := offsetSlice[i].x-offsetSlice[j].x, offsetSlice[i].y-offsetSlice[j].y, offsetSlice[i].z-offsetSlice[j].z

			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			if dz < 0 {
				dz = -dz
			}

			dist := dx + dy + dz
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return maxDist, nil
}
