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

func xxTest_ReturnLeftNodeIfKeyMatches(t *testing.T) {
	root := &binary_tree.TreeNode{Key: 10, Value: 10}
	root.SetLeft(&binary_tree.TreeNode{Key: 5, Value: 15})

	found, _ := binary_tree.Search(5, root)

	assert.Equal(t, 5, found.Key)
}
