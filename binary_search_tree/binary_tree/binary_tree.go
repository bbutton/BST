package binary_tree

type TreeNode struct {
	Key   int
	Value int
	Right *TreeNode
	Left  *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (tn TreeNode) SetLeft(left *TreeNode) {
	tn.Left = left
}

func (tn TreeNode) SetRight(right *TreeNode) {
	tn.Right = right
}

func Search(searchKey int, root *TreeNode) (found *TreeNode, err error) {
	if root == nil {
		return nil, nil
	}

	key := root.Key

	if key == searchKey {
		return root, nil
	}

	return nil, nil
}
