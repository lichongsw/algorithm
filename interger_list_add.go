package main

import (
	"fmt"
)

type listNode struct {
	value int
	next  *listNode
}

func listToValue(node *listNode) int {
	var result int = 0
	currentNode := node
	for currentNode != nil {
		result = result*10 + currentNode.value
		currentNode = currentNode.next
	}

	return result
}

func valueToList(value int) *listNode {
	if value == 0 {
		return nil
	}

	headNode := &listNode{0, nil}

	for value != 0 {
		currentNode := &listNode{0, nil}
		currentNode.value = value % 10
		currentNode.next = headNode.next
		headNode.next = currentNode

		value = value / 10
	}

	return headNode.next
}

func listAdd(nodeA *listNode, nodeB *listNode) *listNode {
	return valueToList(listToValue(nodeA) + listToValue(nodeB))
}

func formatList(node *listNode) (result []int) {
	currentNode := node
	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}
	return result
}

func main() {
	fmt.Println("Welcome to the playground!")

	headNodeA := &listNode{1, nil}
	headNodeA.next = &listNode{2, nil}
	headNodeA.next.next = &listNode{3, nil}
	headNodeA.next.next.next = &listNode{4, nil}

	headNodeB := &listNode{3, nil}
	headNodeB.next = &listNode{4, nil}

	fmt.Println("Value of list A:", listToValue(headNodeA))
	fmt.Println("Items of list A:", formatList(headNodeA))
	fmt.Println("Value of list B:", listToValue(headNodeB))
	fmt.Println("Items of list B:", formatList(headNodeB))

	fmt.Println("Sum of list A and B", formatList(listAdd(headNodeA, headNodeB)))
}
