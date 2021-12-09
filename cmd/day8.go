package cmd

import (
	"errors"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/zap"
)

func init() {
	solvers[8] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay8P1,
		P2: DoDay8P2,
	}

	for i, v := range day8Mapping {
		l := len(v)
		if _, ok := day8LengthMapping[l]; !ok {
			day8LengthMapping[l] = make([]int, 0)
		}

		day8LengthMapping[l] = append(day8LengthMapping[l], i)
	}
}

const (
	a = "a"
	b = "b"
	c = "c"
	d = "d"
	e = "e"
	f = "f"
	g = "g"
)

var day8Mapping = map[int][]string{
	0: {a, b, c, e, f, g},
	1: {c, f},
	2: {a, c, d, e, g},
	3: {a, c, d, f, g},
	4: {b, c, d, f},
	5: {a, b, d, f, g},
	6: {a, b, d, e, f, g},
	7: {a, c, f},
	8: {a, b, c, d, e, f, g},
	9: {a, b, c, d, f, g},
}

var day8LengthMapping = map[int][]int{}

func DoDay8P1(input string) {
	sol, err := solveDay8P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 8 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay8P1(input string) (int, error) {
	lines := strings.Split(input, "\n")
	outputs := make([][]string, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " | ")
		if len(split) != 2 {
			return 0, errors.New("invalid input")
		}

		outputs[i] = strings.Split(split[1], " ")
	}

	lengths := make(map[int]int)
	for _, val := range day8Mapping {
		l := len(val)
		if _, ok := lengths[l]; !ok {
			lengths[l] = 0
		}

		lengths[l]++
	}

	uniques := make([]int, 0, len(lengths))
	for k, v := range lengths {
		if v == 1 {
			uniques = append(uniques, k)
		}
	}

	isUnique := func(s string) bool {
		for _, val := range uniques {
			if len(s) == val {
				return true
			}
		}

		return false
	}

	count := 0

	for _, val := range outputs {
		for _, v := range val {
			if isUnique(v) {
				count++
			}
		}
	}

	return count, nil
}

func DoDay8P2(input string) {
	sol, err := solveDay8P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 8 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay8P2(input string) (int, error) {
	lines := strings.Split(input, "\n")
	outputs := make([][]string, len(lines))
	inputs := make([][]string, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " | ")
		if len(split) != 2 {
			return 0, errors.New("invalid input")
		}

		outputs[i] = strings.Split(split[1], " ")
		inputs[i] = strings.Split(split[0], " ")
	}

	mappings := make(map[string][]string)
	for _, val := range []string{a, b, c, d, e, f, g} {
		mappings[val] = []string{a, b, c, d, e, f, g}
	}

	done := func() bool {
		for _, val := range mappings {
			if len(val) != 1 {
				return false
			}
		}

		return true
	}

	attempts := 0

	for !done() && attempts < 100 {
		attempts++
		for _, input := range inputs {
			chars := strings.Split(input[0], "")
			if func() bool {
				for _, c := range chars {
					if len(mappings[c]) != 1 {
						return false
					}
				}

				return true
			}() {
				things := make([]string, 0)
				for _, c := range chars {
					things = append(things, mappings[c][0])
				}

				num := -1
				for i, val := range day8Mapping {
					if sliceEqual(val, things) {
						num = i
						break
					}
				}

				if num == -1 {
					return 0, errors.New("invalid input")
				}

				mappings = deduce(mappings, num, things)
			}
		}
	}

	zap.L().Info("done", zap.Bool("done", done()))

	return 0, nil
}

func solveDay8P2ForOne(input []string, output []string) map[string]string {
	// letterMapping := make(map[string]string)
	input = append(input, output...)
	letterMapping := make(map[string][]string)
	for _, val := range []string{a, b, c, d, e, f, g} {
		letterMapping[val] = []string{a, b, c, d, e, f, g}
	}

	pivotedLetterMapping := make(map[string][]string)
	for a, bs := range letterMapping {
		for _, b := range bs {
			if _, ok := pivotedLetterMapping[b]; !ok {
				pivotedLetterMapping[b] = []string{}
			}

			pivotedLetterMapping[b] = union(pivotedLetterMapping[b], []string{a})
		}
	}

	unique := func(n int) (bool, []string) {
		letters := day8Mapping[n]
		if len(day8LengthMapping[len(letters)]) == 1 {
			for _, thing := range input {
				if len(thing) == len(letters) {
					return true, strings.Split(thing, "")
				}
			}
			panic("wtf?")
		}

		collection := make([]string, 0)
		for _, thing := range letters {
			collection = union(collection, letterMapping[thing])
		}

		if len(collection) == len(letters) {
			return true, collection
		}

		for _, thing := range input {
			if len(thing) == len(letters) {
				goodCount := 0
				lettersInThing := strings.Split(thing, "")
				collection = make([]string, 0)
				for _, letter := range lettersInThing {
					collection = union(collection, pivotedLetterMapping[letter])
				}

				if !isSubSlice(letters, collection) {
					panic("wtf?")
				}

				for i := 0; i < 10; i++ {
					if isSubSlice(day8Mapping[i], collection) {
						goodCount++
					}
				}

				if goodCount == 1 {
					return true, lettersInThing
				}
			}
		}

		return false, nil
	}

	for count := 0; count < 10; count++ {
		for i := 0; i < 10; i++ {
			if ok, chars := unique(i); ok {
				for _, c := range day8Mapping[i] {
					letterMapping[c] = intersect(letterMapping[c], chars)
				}
			}
		}

		pivotedLetterMapping = make(map[string][]string)
		for a, bs := range letterMapping {
			for _, b := range bs {
				if _, ok := pivotedLetterMapping[b]; !ok {
					pivotedLetterMapping[b] = []string{}
				}

				pivotedLetterMapping[b] = union(pivotedLetterMapping[b], []string{a})
			}
		}
	}

	spew.Dump(letterMapping)
	spew.Dump(pivotedLetterMapping)

	panic("not done")
}

func intersect(a, b []string) []string {
	m := make(map[string]bool)
	n := make([]string, 0)

	for _, val := range a {
		m[val] = true
	}

	for _, val := range b {
		if _, ok := m[val]; ok {
			n = append(n, val)
		}
	}

	return n
}

func union(a, b []string) []string {
	m := make(map[string]bool)
	n := make([]string, 0)

	for _, val := range a {
		m[val] = true
		n = append(n, val)
	}

	for _, val := range b {
		if _, ok := m[val]; !ok {
			n = append(n, val)
		}
	}

	return n
}

func setEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, val := range a {
		if !sliceContains(b, val) {
			return false
		}
	}

	return true
}

func sliceContains(a []string, val string) bool {
	for _, v := range a {
		if v == val {
			return true
		}
	}

	return false
}

func isSubSlice(smol, big []string) bool {
	if len(smol) > len(big) {
		return false
	}

	for _, val := range smol {
		if !sliceContains(big, val) {
			return false
		}
	}

	return true
}

func deduce(m map[string][]string, num int, things []string) map[string][]string {
	targets := day8Mapping[num]
	inTargets := func(s string) bool {
		for _, t := range targets {
			if t == s {
				return true
			}
		}

		return false
	}

	for _, t := range things {
		thingies := make([]string, 0)
		for _, v := range m[t] {
			if inTargets(v) {
				thingies = append(thingies, v)
			}
		}
		m[t] = thingies
	}

	return m
}

func sliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Strings(s1)
	sort.Strings(s2)

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
