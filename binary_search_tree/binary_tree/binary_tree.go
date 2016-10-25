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

func (tn *TreeNode) SetLeft(left *TreeNode) {
	tn.Left = left
}

func (tn *TreeNode) SetRight(right *TreeNode) {
	tn.Right = right
}

func Search(searchKey int, root *TreeNode) (found *TreeNode, err error) {
	if root == nil {
		return nil, nil
	}

	if root.Key == searchKey {
		return root, nil
	} else if searchKey < root.Key {
		return Search(searchKey, root.Left)
	} else if searchKey > root.Key {
		return Search(searchKey, root.Right)
	}

	return nil, nil
}

func traverse(root *TreeNode, traversedNodes []*TreeNode) []*TreeNode {
	if root == nil {
		return traversedNodes
	}

	traversedNodes = traverse(root.Left, traversedNodes)
	traversedNodes = append(traversedNodes, root)
	traversedNodes = traverse(root.Right, traversedNodes)

	return traversedNodes
}

func Traverse(root *TreeNode) (nodes []*TreeNode, err error) {
	traversedNodes := make([]*TreeNode, 0)

	traversedNodes = traverse(root, traversedNodes)

	return traversedNodes, nil
}
