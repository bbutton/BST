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

func (bt BinaryTree) Search(searchKey int) (found *TreeNode, err error) {
	if bt.Root == nil {
		return nil, nil
	}

	key := bt.Root.Key

	if key == searchKey {
		return bt.Root, nil
	}

	return nil, nil
}
