package main

func main() {
	node_44 := TreeNode{
		Val:   44,
		Left:  nil,
		Right: nil,
	}
	node_50 := TreeNode{
		Val:   50,
		Left:  &node_44,
		Right: nil,
	}
	node_58 := TreeNode{
		Val:   58,
		Left:  &node_50,
		Right: nil,
	}
	node_34 := TreeNode{
		Val:   34,
		Left:  nil,
		Right: &node_58,
	}
	node_27 := TreeNode{
		Val:   27,
		Left:  nil,
		Right: &node_34,
	}
	println(minDiffInBST(&node_27))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var min = 2424242424
	var prev = &TreeNode{Val: -1}
	dfs(root, &min, prev)
	return min
}

func dfs(node *TreeNode, min *int, prev *TreeNode) {
	if node == nil {
		return
	}

	dfs(node.Left, min, prev)

	if prev.Val != -1 && node.Val-prev.Val < *min {
		*min = node.Val - prev.Val
	}

	prev.Val = node.Val
	dfs(node.Right, min, prev)
}
