package plus

import (
	"fmt"
)

// Node represents binary tree node.
type Node struct {
	needed bool
	level int
	value int

	parent *Node
	left *Node
	right *Node

	decision string
	controlSum int
}

// AddLevel adds a new level and counts the control sum on it.
func (n *Node) AddLevel(value int) {
	if n.right == nil && n.left == nil {
		n.right = &Node{
			level: n.level + 1,
			value:      value,
			parent:     n,
			left:       nil,
			right:      nil,
			decision:   "+",
			controlSum: n.controlSum + value,
		}
		n.left = &Node{
			level: n.level + 1,
			value: value,
			parent: n,
			left: nil,
			right: nil,
			decision: "-",
			controlSum: n.controlSum - value,
		}
	} else {
		n.left.AddLevel(value)
		n.right.AddLevel(value)
	}
}

// Check checks for the presence of the required control sum in the entire tree.
func (n *Node) Check() bool {
	if n.right == nil && n.left == nil {
		 return n.controlSum == 0
	} else {
		return n.right.Check() || n.left.Check()
	}
}

// Print outputs the entire tree and the characteristics of its nodes.
func (n *Node) Print() {
	if n == nil {
		return
	}
	n.left.Print()
	for i := 0; i < 100 - (n.level * 10); i++ {
		fmt.Print("-")
	}
	fmt.Printf(" %d (%d)[%s%d]\n", n.controlSum, n.level, n.decision, n.value)
	n.right.Print()
}

// FindPath finds all the paths to the needed control sum and writes them to the channel.
func (n *Node) FindPath(resCh chan string) {
	if n.controlSum == 0 && n.right == nil && n.left == nil {
		//go func() {
			tempNode := n
			s := ""
			for tempNode != nil {
				tempNode.needed = true
				s += tempNode.decision
				tempNode = tempNode.parent
			}
			resCh <- s
		//}()
	} else {
		if n.right != nil {
			n.right.FindPath(resCh)
		}
		if n.left != nil {
			n.left.FindPath(resCh)
		}
	}
}