package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay18(t *testing.T) {
	require := require.New(t)

	trivialTree := "[1,2]"
	node, err := day18ParseNode(trivialTree)
	require.NoError(err)
	require.NotNil(node.left)
	require.NotNil(node.right)
	require.Equal(1, *node.left.value)
	require.Equal(2, *node.right.value)
	require.Same(node, node.left.root)
	require.Same(node, node.right.root)
	require.Equal(trivialTree, day18PrintNode(node))

	leftTo2 := day18FindNumberNodeToTheLeft(node.right)
	require.Equal(1, *leftTo2.value)
	rightTo2 := day18FindNumberNodeToTheRight(node.right)
	require.Nil(rightTo2)
	rightTo1 := day18FindNumberNodeToTheRight(node.left)
	require.Equal(2, *rightTo1.value)

	lessTrivialTree := "[9,[8,7]]"
	node, err = day18ParseNode(lessTrivialTree)
	require.NoError(err)
	require.NotNil(node.left)
	require.NotNil(node.right)
	require.Equal(9, *node.left.value)
	require.NotNil(node.right.left)
	require.NotNil(node.right.right)
	require.Equal(8, *node.right.left.value)
	require.Equal(7, *node.right.right.value)

	require.Equal(lessTrivialTree, day18PrintNode(node))

	leftTo7 := day18FindNumberNodeToTheLeft(node.right.right)
	require.Equal(8, *leftTo7.value)
	leftTo8 := day18FindNumberNodeToTheLeft(node.right.left)
	require.Equal(9, *leftTo8.value)
	leftTo9 := day18FindNumberNodeToTheLeft(node.left)
	require.Nil(leftTo9)

	treeToExplode, err := day18ParseNode("[[[[[9,8],1],2],3],4]")
	require.NoError(err)
	exploded := day18MaybeExplodeNode(treeToExplode, 0)
	require.True(exploded)

	t.Log(day18PrintNode(treeToExplode))

	require.NotNil(treeToExplode.left.left.left.left.value)
	require.Equal(0, *treeToExplode.left.left.left.left.value)

	treeToSplit, err := day18ParseNode("[1,15]")
	require.NoError(err)
	splitted := day18MaybeSplitNode(treeToSplit)
	require.True(splitted)
	require.Equal("[1,[7,8]]", day18PrintNode(treeToSplit))

	treeToReduce, err := day18ParseNode("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]")
	require.NoError(err)
	day18ReduceNode(treeToReduce)
	require.Equal("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", day18PrintNode(treeToReduce))

	treeToMagnitude, err := day18ParseNode("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	require.NoError(err)
	require.Equal(1384, day18NodeMagnitude(treeToMagnitude))

	testInput := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	p1, err := solveDay18P1(testInput)
	require.NoError(err)
	require.Equal(4140, p1)

	p2, err := solveDay18P2(testInput)
	require.NoError(err)
	require.Equal(3993, p2)
}
