package main

func main() {
	node_1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node_3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node_6 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node_9 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	node_2 := TreeNode{
		Val:   2,
		Left:  &node_1,
		Right: &node_3,
	}
	node_7 := TreeNode{
		Val:   7,
		Left:  &node_6,
		Right: &node_9,
	}
	node_4 := TreeNode{
		Val:   4,
		Left:  &node_2,
		Right: &node_7,
	}

	res := invertTree(&node_4)
	println(res.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left

	return root
}
