package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func init() {
	solvers[18] = struct {
		P1 func(string)
		P2 func(string)
	}{
		P1: DoDay18P1,
		P2: DoDay18P2,
	}
}

type day18Node struct {
	left  *day18Node
	right *day18Node
	root  *day18Node
	value *int
}

func DoDay18P1(input string) {
	sol, err := solveDay18P1(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 18 part 1", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay18P1(input string) (int, error) {
	lines := strings.Split(input, "\n")

	sum, err := day18ParseNode(lines[0])
	if err != nil {
		return 0, err
	}

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		entry, err := day18ParseNode(line)
		if err != nil {
			return 0, err
		}

		sum, err = day18NodeAddition(sum, entry)
		if err != nil {
			return 0, err
		}

		day18ReduceNode(sum)
	}

	return day18NodeMagnitude(sum), nil
}

func DoDay18P2(input string) {
	sol, err := solveDay18P2(input)
	if err != nil {
		zap.L().Fatal("Failed to solve day 18 part 2", zap.Error(err))
	}

	zap.L().Info("result", zap.Int("solution", sol))
}

func solveDay18P2(input string) (int, error) {
	max := 0

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			one, err := day18ParseNode(lines[i])
			if err != nil {
				return 0, err
			}

			two, err := day18ParseNode(lines[j])
			if err != nil {
				return 0, err
			}

			sum, err := day18NodeAddition(one, two)
			if err != nil {
				return 0, err
			}

			day18ReduceNode(sum)
			if mag := day18NodeMagnitude(sum); mag > max {
				max = mag
			}

			one, err = day18ParseNode(lines[i])
			if err != nil {
				return 0, err
			}

			two, err = day18ParseNode(lines[j])
			if err != nil {
				return 0, err
			}

			sum, err = day18NodeAddition(two, one)
			if err != nil {
				return 0, err
			}

			day18ReduceNode(sum)
			if mag := day18NodeMagnitude(sum); mag > max {
				max = mag
			}
		}
	}

	return max, nil
}

func day18ParseNode(input string) (*day18Node, error) {
	if input[0] != '[' {
		// we're at a number node
		num, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}

		return &day18Node{
			value: &num,
		}, nil
	}

	// we're at a tree node
	node := &day18Node{}
	input = input[1 : len(input)-1]
	level := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			level++
		} else if input[i] == ']' {
			level--
		} else if input[i] == ',' && level == 0 {
			left, err := day18ParseNode(input[:i])
			if err != nil {
				return nil, err
			}

			right, err := day18ParseNode(input[i+1:])
			if err != nil {
				return nil, err
			}

			left.root = node
			right.root = node
			node.left = left
			node.right = right

			return node, nil
		}
	}

	return nil, errors.New("invalid input")
}

func day18ReduceNode(node *day18Node) {
	for {
		if day18MaybeExplodeNode(node, 0) {
			continue
		}

		if day18MaybeSplitNode(node) {
			continue
		}

		break
	}
}

func day18MaybeExplodeNode(node *day18Node, depth int) bool {
	if node.left == nil && node.right == nil {
		return false
	}

	if node.left.value != nil && node.right.value != nil {
		// we're at a number node
		if depth == 4 {
			// we need to explode
			leftNumNode := day18FindNumberNodeToTheLeft(node)
			if leftNumNode != nil {
				newVal := *leftNumNode.value + *node.left.value
				leftNumNode.value = &newVal
			}

			rightNumNode := day18FindNumberNodeToTheRight(node)
			if rightNumNode != nil {
				newVal := *rightNumNode.value + *node.right.value
				rightNumNode.value = &newVal
			}

			root := node.root
			if root == nil {
				panic("wtf?")
			}

			zero := 0

			if root.left == node {
				root.left = &day18Node{
					value: &zero,
					root:  root,
				}
			} else if root.right == node {
				root.right = &day18Node{
					value: &zero,
					root:  root,
				}
			} else {
				panic("wtf?")
			}

			return true
		}
		return false
	}

	if node.left != nil {
		if day18MaybeExplodeNode(node.left, depth+1) {
			return true
		}
	}

	if node.right != nil {
		if day18MaybeExplodeNode(node.right, depth+1) {
			return true
		}
	}

	return false
}

func day18MaybeSplitNode(node *day18Node) bool {
	if node.left != nil {
		if day18MaybeSplitNode(node.left) {
			return true
		}
	}

	if node.right != nil {
		if day18MaybeSplitNode(node.right) {
			return true
		}
	}

	if node.left != nil || node.right != nil {
		return false
	}

	// we're in a number node
	if node.value == nil {
		panic("wtf?")
	}

	value := *node.value
	if value >= 10 {
		leftNum := value / 2
		rightNum := value - leftNum

		leftNode := &day18Node{
			value: &leftNum,
			root:  node,
		}
		rightNode := &day18Node{
			value: &rightNum,
			root:  node,
		}

		node.left = leftNode
		node.right = rightNode
		node.value = nil

		return true
	}

	return false
}

func day18FindNumberNodeToTheLeft(node *day18Node) *day18Node {
	if node.root == nil {
		return nil
	}
	root := node.root
	if root.right == node {
		// we're a right node
		return day18FindRightMostNode(root.left)
	} else {
		// we're a left node
		return day18FindNumberNodeToTheLeft(root)
	}
}

func day18FindNumberNodeToTheRight(node *day18Node) *day18Node {
	if node.root == nil {
		return nil
	}
	root := node.root
	if root.left == node {
		// we're a left node
		return day18FindLeftMostNode(root.right)
	} else {
		// we're a right node
		return day18FindNumberNodeToTheRight(root)
	}
}

func day18FindRightMostNode(node *day18Node) *day18Node {
	if node.right == nil {
		return node
	}

	return day18FindRightMostNode(node.right)
}

func day18FindLeftMostNode(node *day18Node) *day18Node {
	if node.left == nil {
		return node
	}

	return day18FindLeftMostNode(node.left)
}

func day18PrintNode(node *day18Node) string {
	if node.left == nil && node.right == nil {
		return fmt.Sprintf("%d", *node.value)
	}

	return fmt.Sprintf("[%s,%s]", day18PrintNode(node.left), day18PrintNode(node.right))
}

func day18NodeAddition(left, right *day18Node) (*day18Node, error) {
	if left.root != nil || right.root != nil {
		return nil, errors.New("can only add root nodes")
	}

	newRoot := &day18Node{
		left:  left,
		right: right,
	}

	left.root = newRoot
	right.root = newRoot

	return newRoot, nil
}

func day18NodeMagnitude(node *day18Node) int {
	if node.left == nil && node.right == nil {
		return *node.value
	}

	if node.left == nil || node.right == nil {
		panic("wtf?")
	}

	return 3*day18NodeMagnitude(node.left) + 2*day18NodeMagnitude(node.right)
}
