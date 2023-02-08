package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1
	l1 := ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := ListNode{
		Val:  4,
		Next: &l1,
	}
	l3 := ListNode{
		Val:  3,
		Next: &l2,
	}

	// 2
	l4 := ListNode{
		Val:  5,
		Next: nil,
	}

	l5 := ListNode{
		Val:  6,
		Next: &l4,
	}

	l6 := ListNode{
		Val:  4,
		Next: &l5,
	}

	res := addTwoNumbers(&l6, &l3)
	println(&res.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumList := &ListNode{}
	var currNode *ListNode
	currNode = sumList

	x, y, curry, sum := 0, 0, 0, 0
	for l1 != nil || l2 != nil || curry != 0 {
		x, y = 0, 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum = x + y + curry
		curry = sum / 10

		newNode := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		currNode.Next = newNode
		currNode = currNode.Next
	}

	return sumList.Next
}
