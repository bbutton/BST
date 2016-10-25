package binary_tree_test

import (
	"testing"

	"github.com/bbutton/BST/binary_search_tree/binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestSearchingEmptyTreeAlwaysReturnsNil(t *testing.T) {
	result, _ := binary_tree.Search(1, nil)

	assert.Nil(t, result)
}

func Test_RootNodeReturnedIfMatches(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 1, Value: 10}

	result, _ := binary_tree.Search(1, root)

	assert.Equal(t, result.Key, 1)
	assert.Equal(t, result.Value, 10)
}

func Test_NilReturnedIfKeyDoesNotMatch(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 1, Value: 10}

	found, _ := binary_tree.Search(15, root)

	assert.Nil(t, found)
}

func Test_ReturnLeftNodeIfKeyMatches(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	root.SetLeft(&binary_tree.TreeNode{Key: 5, Value: 15})

	assert.NotNil(t, root.Left)

	found, _ := binary_tree.Search(5, root)

	assert.NotNil(t, found)
	assert.Equal(t, 5, found.Key)
}

func Test_WillRecurseSeveralLevelsOfLeftToMatchKey(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	firstLeft := &binary_tree.TreeNode{Key: 5, Value: 15}
	secondLeft := &binary_tree.TreeNode{Key: 4, Value: 15}

	firstLeft.SetLeft(secondLeft)
	root.SetLeft(firstLeft)

	found, _ := binary_tree.Search(4, root)

	assert.Equal(t, 4, found.Key)
}

func Test_WillSearchRightWhenSearchKeyGreaterThanNodeKey(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	root.SetRight(&binary_tree.TreeNode{Key: 15, Value: 15})

	found, _ := binary_tree.Search(15, root)

	assert.NotNil(t, found)
	assert.Equal(t, 15, found.Key)
}

func Test_TraversingEmptyTreeReturnsEmptySlice(t *testing.T) {
	traversedNodes, _ := binary_tree.Traverse(nil)

	assert.Equal(t, 0, len(traversedNodes))
}

func Test_TraversingSingleNodeTreeReturnsOneNode(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 1, Value: 10}

	traversedNodes, _ := binary_tree.Traverse(root)

	assert.Equal(t, 1, traversedNodes[0].Key)
}

func Test_TraversingLeftSideReturnsInIncreasingOrder(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	left := &binary_tree.TreeNode{Key: 5, Value: 5}
	root.Left = left

	traversedNodes, _ := binary_tree.Traverse(root)

	assert.Equal(t, 5, traversedNodes[0].Key)
	assert.Equal(t, 10, traversedNodes[1].Key)
}

func Test_TraversingRightSideReturnsInIncreasingOrder(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	right := &binary_tree.TreeNode{Key: 15, Value: 5}
	root.Right = right

	traversedNodes, _ := binary_tree.Traverse(root)

	assert.Equal(t, 10, traversedNodes[0].Key)
	assert.Equal(t, 15, traversedNodes[1].Key)
}

/*
func Test_MoreComplicatedTreeShouldReturnInCorrectOrder(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	l1 := &binary_tree.TreeNode{Key: 5, Value: 5}
	l2 := &binary_tree.TreeNode{Key: 1, Value: 1}
	l3 := &binary_tree.TreeNode{Key: 7, Value: 3}
	l4 := &binary_tree.TreeNode{Key: 6, Value: 2}
	root.Left = l1
	l1.Left = l2
	l1.Right = l3
	l3.Left = l4

	r1 := &binary_tree.TreeNode{Key: 15, Value: 15}
	r2 := &binary_tree.TreeNode{Key: 12, Value: 12}
	r1.Left = r2

	r3 := &binary_tree.TreeNode{Key: 11, Value: 11}
	r2.Left = r3

	r4 := &binary_tree.TreeNode{Key: 13, Value: 13}
	r2.Right = r4

	r5 := &binary_tree.TreeNode{Key: 20, Value: 20}
	r1.Right = r5

	r6 := &binary_tree.TreeNode{Key: 17, Value: 17}
	r5.Left = r6

	r7 := &binary_tree.TreeNode{Key: 16, Value: 16}
	r6.Left = r7

	r8 := &binary_tree.TreeNode{Key: 18, Value: 18}
	r6.Right = r8

	r9 := &binary_tree.TreeNode{Key: 21, Value: 21}
	r5.Right = r9

	root.Right = r1

	traversedNodes, _ := binary_tree.Traverse(root)

	for i := 0; i < len(traversedNodes); i++ {
		fmt.Printf("%d: %d\n", i, traversedNodes[i].Key)
	}
}
*/
