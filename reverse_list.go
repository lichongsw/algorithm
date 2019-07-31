package main

import (
	"fmt"
)

type listNode struct {
	value int
	next  *listNode
}

func reverseList(list *listNode) *listNode {
	var previousNode *listNode = nil
	currentNode := list
	for {
		if currentNode != nil {
			nextNode := currentNode.next
			currentNode.next = previousNode
			previousNode = currentNode
			currentNode = nextNode
		} else {
			break
		}
	}

	return previousNode
}

func main() {
	fmt.Println("Welcome to the playground!")
	nodeA := &listNode{10, nil}
	nodeB := &listNode{11, nil}
	nodeC := &listNode{12, nil}

	nodeA.next = nodeB
	nodeB.next = nodeC
	nodeC.next = nil

	fmt.Println("Result value for node C:", reverseList(nodeA).value)
}
