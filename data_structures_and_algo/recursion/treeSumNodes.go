package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	root := &Node{data: 2}
	root.left = &Node{data: 3}
	root.right = &Node{data: 4}
	root.left.left = &Node{data: 5}
	root.left.right = &Node{data: 6}
	println(sumNodes(root))
}

func sumNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return node.data + sumNodes(node.left) + sumNodes(node.right)
}
