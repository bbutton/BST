package binary_tree_test

import (
	"testing"

	"github.com/bbutton/BST/binary_search_tree/binary_tree"
	"github.com/stretchr/testify/assert"
)

func TestSearchingEmptyTreeAlwaysReturnsNil(t *testing.T) {
	binaryTree := binary_tree.BinaryTree{}

	result, _ := binaryTree.Search(1)

	assert.Nil(t, result)
}

func Test_RootNodeReturnedIfMatches(t *testing.T) {
	binaryTree := binary_tree.BinaryTree{Root: &binary_tree.TreeNode{Key: 1, Value: 10}}

	found, _ := binaryTree.Search(1)

	assert.Equal(t, found.Key, 1)
	assert.Equal(t, found.Value, 10)
}

func Test_NilReturnedIfKeyDoesNotMatch(t *testing.T) {
	binaryTree := binary_tree.BinaryTree{Root: &binary_tree.TreeNode{Key: 1, Value: 10}}

	found, _ := binaryTree.Search(15)

	assert.Nil(t, found)
}
