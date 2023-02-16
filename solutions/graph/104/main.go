package main

func main() {
	node_7 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node_15 := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	node_20 := TreeNode{
		Val:   20,
		Left:  &node_15,
		Right: &node_7,
	}
	node_9 := TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	node_3 := TreeNode{
		Val:   3,
		Left:  &node_9,
		Right: &node_20,
	}
	println(maxDepth(&node_3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	length, l, r := 1, 0, 0
	if root.Left == nil && root.Right == nil {
		return length
	}
	if root.Left != nil {
		l = maxDepth(root.Left)
	}
	if root.Right != nil {
		r = maxDepth(root.Right)
	}

	if l > r {
		length += l
	} else {
		length += r
	}

	return length
}
