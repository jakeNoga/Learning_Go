package main

import "fmt"

type Node[C comparable] struct {
	Val      C
	NextNode *Node[C]
}

func (N *Node[C]) Add(value C) {
	thisNode := &Node[C]{
		Val:      value,
		NextNode: nil,
	}
	fmt.Println("Value:", value)
	if N == nil {
		fmt.Println("This should have added it here")
		N = thisNode
		fmt.Println(N)
	}
	for curNode := N; curNode != nil; curNode = curNode.NextNode {
		if curNode.NextNode == nil {
			curNode.NextNode = thisNode
			fmt.Println("Current Node =", curNode)
			fmt.Println("This node", curNode.NextNode)
			break
		}
	}

}

func (N *Node[C]) insert(value C, pos int) {
	if N == nil {
		fmt.Println("Unable to add the value to the specified index.")
		return
	}

	fmt.Println("value", value)
	fmt.Println("pos", pos)
	thisNode := &Node[C]{
		Val: value,
	}

	leftNode := N
	for i := 1; i < pos; i++ {
		fmt.Println("i:", i)
		if leftNode.NextNode != nil {
			leftNode = leftNode.NextNode
		} else {
			break
		}
	}

	fmt.Println("leftNode", leftNode)
	rightNode := leftNode.NextNode

	fmt.Println("rightNode:", rightNode)
	thisNode.NextNode = rightNode
	leftNode.NextNode = thisNode
}

func (N *Node[C]) index(value C) int {
	tempNode := N
	for i := 0; tempNode != nil; i++ {
		if tempNode.Val == value {
			return i
		}
		tempNode = tempNode.NextNode
	}
	return -1
}

func main() {
	head := &Node[int]{
		Val: 1,
	}
	// fmt.Println(head)
	// head.Add(1)
	fmt.Println(head)
	head.Add(10)
	fmt.Println(head)
	head.Add(22)
	fmt.Println("Starting index's")
	fmt.Println(head.index(1))
	fmt.Println(head.index(10))
	fmt.Println(head.index(22))

	fmt.Println("Starting inserts")
	head.insert(5, 1)
	head.insert(8, 2)
	fmt.Println("Starting second index's")
	fmt.Println(head.index(1))
	fmt.Println(head.index(5))
	fmt.Println(head.index(8))
	fmt.Println(head.index(10))
	fmt.Println(head.index(22))
}
