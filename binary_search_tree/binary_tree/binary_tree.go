package binary_tree

import "fmt"

type TreeNode struct {
	Key   int
	Value int
	Right *TreeNode
	Left  *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (tn *TreeNode) SetLeft(left *TreeNode) {
	tn.Left = left
}

func (tn *TreeNode) SetRight(right *TreeNode) {
	tn.Right = right
}

func Search(searchKey int, root *TreeNode) (found *TreeNode, err error) {
	fmt.Println("Entering search")

	if root == nil {
		fmt.Println("Leaving search, root == nil")
		return nil, nil
	}

	fmt.Printf("Key: %d, Searchkey: %d\n", root.Key, searchKey)

	if root.Key == searchKey {
		fmt.Println("Returning match on root")
		return root, nil
	} else if searchKey < root.Key {
		fmt.Println("Looking for match on Left")
		return Search(searchKey, root.Left)
	} else if searchKey > root.Key {
		fmt.Println("Looking for march on Right")
		return Search(searchKey, root.Right)
	}

	fmt.Println("No match at all, returning nil, nil")
	return nil, nil
}
