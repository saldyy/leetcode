package problem0226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(temp)

	return root
}

func Execute(a *TreeNode) *TreeNode {

	return invertTree(a)
}
