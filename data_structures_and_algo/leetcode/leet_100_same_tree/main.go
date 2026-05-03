package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	return isSameTree(p.Right, q.Right)
}

func main() {
	pNode3 := &TreeNode{Val: 3}
	pNode2 := &TreeNode{Val: 2}
	pNode1 := &TreeNode{Val: 1, Left: pNode3, Right: pNode2}

	qNode3 := &TreeNode{Val: 3}
	qNode2 := &TreeNode{Val: 2}
	qNode1 := &TreeNode{Val: 1, Left: qNode3, Right: qNode2}

	fmt.Println(isSameTree(pNode1, qNode1))
}
