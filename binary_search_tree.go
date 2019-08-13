package main

import (
	"fmt"
)

type BinarytreeNode struct {
	Value int
	Left  *BinarytreeNode
	Right *BinarytreeNode
}

func getMaxNode(node *BinarytreeNode) *BinarytreeNode {
	if node == nil {
		return nil
	}

	currentNode := node
	for currentNode.Right != nil {
		currentNode = node.Right
	}

	return currentNode
}

func InsertNode(node *BinarytreeNode, value int) *BinarytreeNode {
	if node == nil {
		return &BinarytreeNode{value, nil, nil}
	}

	if value < node.Value {
		node.Left = InsertNode(node.Left, value)
	} else {
		node.Right = InsertNode(node.Right, value)
	}

	return node
}

func SearchNode(node *BinarytreeNode, value int) (foundNode *BinarytreeNode, parentNode *BinarytreeNode) {
	if node == nil {
		return nil, nil
	}

	currentNode := node
	for currentNode != nil && currentNode.Value != value {
		parentNode = currentNode
		if parentNode.Value > value {
			currentNode = parentNode.Left
		} else {
			currentNode = parentNode.Right
		}
	}

	return currentNode, parentNode
}

func DeleteNode(node *BinarytreeNode, value int) {
	if node == nil {
		fmt.Println("Delete nil node!")
	}

	foundNode, parentNode := SearchNode(node, value)
	if foundNode == nil {
		fmt.Println("Node to delete is not found!")
	}
	if parentNode == nil {
		// delete root node
		foundNode = nil
	}

	// leaf node without child
	if foundNode.Left == nil && foundNode.Right == nil {
		if parentNode.Left == foundNode {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
		foundNode = nil
		return
	}

	// left child only
	if foundNode.Left != nil && foundNode.Right == nil {
		if parentNode.Left == foundNode {
			parentNode.Left = foundNode.Left
		} else {
			parentNode.Right = foundNode.Left
		}
		foundNode = nil
		return
	}

	// right child only
	if foundNode.Left == nil && foundNode.Right != nil {
		if parentNode.Left == foundNode {
			parentNode.Left = foundNode.Right
		} else {
			parentNode.Right = foundNode.Right
		}
		foundNode = nil
		return
	}

	// both left and right children
	if foundNode.Left != nil && foundNode.Right != nil {
		exchangeNode := getMaxNode(foundNode.Left)
		// fmt.Println("Exchange node value: ", exchangeNode.Value)
		DeleteNode(node, exchangeNode.Value)
		foundNode.Value = exchangeNode.Value
		return
	}
}

func PreOrderSlice(node *BinarytreeNode) (result []int) {
	if node != nil {
		result = append(result, node.Value)
		result = append(result, PreOrderSlice(node.Left)...)
		result = append(result, PreOrderSlice(node.Right)...)
	}

	return result
}

func InOrderSlice(node *BinarytreeNode) (result []int) {
	if node != nil {
		result = append(result, InOrderSlice(node.Left)...)
		result = append(result, node.Value)
		result = append(result, InOrderSlice(node.Right)...)
	}

	return result
}

func PostOrderSlice(node *BinarytreeNode) (result []int) {
	if node != nil {
		result = append(result, PostOrderSlice(node.Left)...)
		result = append(result, PostOrderSlice(node.Right)...)
		result = append(result, node.Value)
	}

	return result
}

func GetLevel(node *BinarytreeNode) (result int) {
	if node == nil {
		return 0
	}

	level := Max(GetLevel(node.Left), GetLevel(node.Right)) + 1
	return level
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println("Welcome to the playground!")

	rootNode := &BinarytreeNode{10, nil, nil}
	InsertNode(rootNode, 6)
	InsertNode(rootNode, 15)
	InsertNode(rootNode, 4)
	InsertNode(rootNode, 8)
	InsertNode(rootNode, 12)
	InsertNode(rootNode, 18)
	InsertNode(rootNode, 5)
	InsertNode(rootNode, 7)
	InsertNode(rootNode, 9)
	InsertNode(rootNode, 16)
	InsertNode(rootNode, 19)

	fmt.Println("Levels of tree: ", GetLevel(rootNode))

	fmt.Println("Pre order print: ", PreOrderSlice(rootNode))
	fmt.Println("In order print: ", InOrderSlice(rootNode))
	fmt.Println("Post order print: ", PostOrderSlice(rootNode))

	foundNode, parentNode := SearchNode(rootNode, 5)
	fmt.Println("Search node result: ", foundNode.Value, ",parent node value", parentNode.Value)

	foundNode, parentNode = SearchNode(rootNode, 18)
	fmt.Println("Search node result: ", foundNode.Value, ",parent node value", parentNode.Value)

	DeleteNode(rootNode, 15)
	foundNode, parentNode = SearchNode(rootNode, 18)
	fmt.Println("Search node result: ", foundNode.Value, ",parent node value", parentNode.Value)
	InsertNode(rootNode, 15)
	foundNode, parentNode = SearchNode(rootNode, 18)
	fmt.Println("Search node result: ", foundNode.Value, ",parent node value", parentNode.Value)

	fmt.Println("Pre order print: ", PreOrderSlice(rootNode))
	fmt.Println("In order print: ", InOrderSlice(rootNode))
	fmt.Println("Post order print: ", PostOrderSlice(rootNode))
}
