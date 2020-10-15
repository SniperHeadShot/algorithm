package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	recursionBFS([]*Node{root})
	return root
}

func recursionBFS(nodes []*Node) {
	nodeLen := len(nodes)
	if nodeLen == 0 {
		return
	}
	childNodes := make([]*Node, 0, 0)
	for i, n := range nodes {
		if n == nil {
			break
		}
		if i+1 == nodeLen {
			nodes[i].Next = nil
		} else {
			nodes[i].Next = nodes[i+1]
		}
		childNodes = append(childNodes, n.Left, n.Right)
	}
	recursionBFS(childNodes)
}
