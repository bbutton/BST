package main

import "github.com/bbutton/BST/binary_search_tree/binary_tree"

func main() {
	rootNode := binary_tree.TreeNode{Key: 10, Value: 15}
	binaryTree := binary_tree.BinaryTree{Root: &rootNode}

	l1 := binary_tree.TreeNode{Key: 5, Value: 10}
	r1 := binary_tree.TreeNode{Key: 15, Value: 4}
	rootNode.SetLeft(&l1)
	rootNode.SetRight(&r1)

	binaryTree.Search(15)
}
