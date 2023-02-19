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
	println(zigzagLevelOrder(&node_3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	ans = zigzag(root, 0, ans)
	return ans
}

func zigzag(node *TreeNode, lvl int, ans [][]int) [][]int {
	if node == nil {
		return ans
	}

	if len(ans) <= lvl {
		ans = append(ans, []int{node.Val})
	} else if lvl%2 == 0 {
		ans[lvl] = append(ans[lvl], node.Val)
	} else {
		ans[lvl] = append([]int{node.Val}, ans[lvl]...)
	}

	lvl++
	ans = zigzag(node.Left, lvl, ans)
	ans = zigzag(node.Right, lvl, ans)
	return ans
}
