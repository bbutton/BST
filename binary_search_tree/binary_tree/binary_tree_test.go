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
