package cmd

import (
	"errors"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[12] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay12P1,
		P2: DoDay12P2,
	}
}

type day12Node struct {
	name      string
	connected []*day12Node
}

func day12ParseGraph(input string) (map[string]*day12Node, error) {
	nodes := make(map[string]*day12Node)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			return nil, errors.New("invalid line")
		}

		if _, ok := nodes[split[0]]; !ok {
			nodes[split[0]] = &day12Node{
				name:      split[0],
				connected: make([]*day12Node, 0),
			}
		}

		if _, ok := nodes[split[1]]; !ok {
			nodes[split[1]] = &day12Node{
				name:      split[1],
				connected: make([]*day12Node, 0),
			}
		}

		if !day12NodeInList(nodes[split[1]], nodes[split[0]].connected) {
			nodes[split[0]].connected = append(nodes[split[0]].connected, nodes[split[1]])
		}

		if !day12NodeInList(nodes[split[0]], nodes[split[1]].connected) {
			nodes[split[1]].connected = append(nodes[split[1]].connected, nodes[split[0]])
		}
	}

	return nodes, nil
}

func day12NodeInList(node *day12Node, list []*day12Node) bool {
	for _, n := range list {
		if n.name == node.name {
			return true
		}
	}

	return false
}

func day12P2CanAppend(node *day12Node, list []*day12Node) bool {
	lcMap := make(map[string]int)
	for _, n := range list {
		if n.name != "start" && n.name != "end" && strings.ToLower(n.name) == n.name {
			if _, ok := lcMap[n.name]; !ok {
				lcMap[n.name] = 1
			} else {
				lcMap[n.name]++
			}
		}
	}

	dupExists := false

	for _, v := range lcMap {
		if v >= 2 {
			dupExists = true
		}
	}

	if dupExists {
		_, ok := lcMap[node.name]
		return !ok
	} else {
		count, ok := lcMap[node.name]
		if !ok {
			return true
		}
		return count == 1
	}
}

func DoDay12P1(input string) {
	sol, err := solveDay12P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 12 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay12P1(input string) (int, error) {
	graph, err := day12ParseGraph(input)
	if err != nil {
		return 0, err
	}

	type pool struct {
		nodes    []*day12Node
		terminal bool
		valid    bool
	}

	workPool := make([]pool, 0)

	workPool = append(workPool, pool{
		nodes:    []*day12Node{graph["start"]},
		terminal: false,
		valid:    true,
	})

	done := func() bool {
		for _, p := range workPool {
			if !p.terminal || p.valid {
				return false
			}
		}

		return true
	}

	count := 0

	path := make([][]*day12Node, 0)

	for !done() && count < 1000 {
		count++
		newPool := make([]pool, 0)
		for _, vLoop := range workPool {
			v := pool{
				terminal: vLoop.terminal,
				nodes:    make([]*day12Node, len(vLoop.nodes)),
				valid:    vLoop.valid,
			}

			for i, n := range vLoop.nodes {
				val := *n
				ptr := &val
				v.nodes[i] = ptr
			}

			if v.terminal {
				continue
			}

			if !v.valid {
				continue
			}

			for _, c := range v.nodes[len(v.nodes)-1].connected {
				c := c
				if strings.ToLower(c.name) != c.name {
					// we're upper case
					newPool = append(newPool, pool{
						nodes:    append(v.nodes, c),
						valid:    true,
						terminal: false,
					})
				} else {
					// we're lower case
					if c.name == "end" {
						path = append(path, append(v.nodes, c))
					} else {
						if !day12NodeInList(c, v.nodes) {
							newNodes := append(v.nodes, c)
							newPool = append(newPool, pool{
								nodes:    newNodes,
								valid:    true,
								terminal: false,
							})
						}
					}
				}
			}
		}

		workPool = newPool

	}

	if !done() {
		return 0, errors.New("hit 1000 iterations")
	}

	return len(path), nil
}

func DoDay12P2(input string) {
	sol, err := solveDay12P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 12 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay12P2(input string) (int, error) {
	graph, err := day12ParseGraph(input)
	if err != nil {
		return 0, err
	}

	type pool struct {
		nodes    []*day12Node
		terminal bool
		valid    bool
	}

	workPool := make([]pool, 0)

	workPool = append(workPool, pool{
		nodes:    []*day12Node{graph["start"]},
		terminal: false,
		valid:    true,
	})

	done := func() bool {
		for _, p := range workPool {
			if !p.terminal || p.valid {
				return false
			}
		}

		return true
	}

	count := 0

	path := make([][]*day12Node, 0)

	for !done() && count < 1000 {
		count++
		newPool := make([]pool, 0)
		for _, vLoop := range workPool {
			v := pool{
				terminal: vLoop.terminal,
				nodes:    make([]*day12Node, len(vLoop.nodes)),
				valid:    vLoop.valid,
			}

			for i, n := range vLoop.nodes {
				val := *n
				ptr := &val
				v.nodes[i] = ptr
			}

			if v.terminal {
				continue
			}

			if !v.valid {
				continue
			}

			for _, c := range v.nodes[len(v.nodes)-1].connected {
				c := c
				if strings.ToLower(c.name) != c.name {
					// we're upper case
					newPool = append(newPool, pool{
						nodes:    append(v.nodes, c),
						valid:    true,
						terminal: false,
					})
				} else {
					// we're lower case
					if c.name == "end" {
						path = append(path, append(v.nodes, c))
					} else if c.name == "start" {
						continue
					} else {
						if day12P2CanAppend(c, v.nodes) {
							newNodes := append(v.nodes, c)
							newPool = append(newPool, pool{
								nodes:    newNodes,
								valid:    true,
								terminal: false,
							})
						}
					}
				}
			}
		}

		workPool = newPool
	}

	if !done() {
		return 0, errors.New("hit 1000 iterations")
	}

	return len(path), nil
}
