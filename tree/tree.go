package tree

import "fmt"

type BinaryTree struct {
	root *BinaryNode
}

type BinaryNode struct {
	data  int64
	left  *BinaryNode
	right *BinaryNode
}

func (t *BinaryTree) Insert(data int64) {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
}

func (t *BinaryTree) ListAll(printType string) {

	switch printType {
	case "dfs":
		t.root.listDFS()
		break
	case "bfs":
		t.root.listBFS()
		break

	default:
		fmt.Printf("Enter type of traversal please")
		return
	}

}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *BinaryNode) listDFS() {
	if n == nil {
		return
	}

	fmt.Printf("%d,", n.data)

	if n.left != nil {
		n.left.listDFS()
	}

	if n.right != nil {
		n.right.listDFS()
	}
}

func (n *BinaryNode) listBFS() {

	queue := []*BinaryNode{n}

	for len(queue) > 0 {
		lastElement := queue[0]

		fmt.Printf("%d,", lastElement.data)

		if lastElement.left != nil {
			queue = append(queue, lastElement.left)
		}

		if lastElement.right != nil {
			queue = append(queue, lastElement.right)
		}

		queue = queue[1:]
	}

}
