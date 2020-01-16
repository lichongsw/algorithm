package main

import (
	"fmt"
	"strconv"
)

// test tree for serialize and deserialize, same value of different nodes are also considered
//         1
//    |    |    |
//    3    2    2
//   | |   |    
//   5 6   8    

type NArrayNode struct {
	value     int
	childrens []*NArrayNode
}

// use two kind of flag to indicate the node itself "N" and patent "P", the parent of the root use special "R" flag
func Serialize(root *NArrayNode) string {
	if root == nil {
		return ""
	}

	result := ""
	nodeQueue := make([]*NArrayNode, 1)
	nodeQueue[0] = root
	result += "N" + strconv.Itoa(root.value) + "P" + "R"
	for len(nodeQueue) > 0 {
		for _, child := range nodeQueue[0].childrens {
			nodeQueue = append(nodeQueue, child)
			result += "N" + strconv.Itoa(child.value) + "P" + strconv.Itoa(nodeQueue[0].value)
		}

		nodeQueue = nodeQueue[1:]
	}

	return result
}

func DeSerialize(serializedString string) *NArrayNode {
	// the serialize is based on level, so the parent is always created before child
	root := &NArrayNode{}
	nodeMap := make(map[int]*NArrayNode, 0)
	for i:=0; i<len(serializedString);i++ {
		if serializedString[i] == 'N' {
			for j:=i+1; j<len(serializedString); j++ {
				// get node one by one
				if serializedString[j] == 'P' {
					// create node itself
					selfValue,_ := strconv.Atoi(string(serializedString[i+1:j]))
					selfNode := &NArrayNode{selfValue, nil}
					nodeMap[selfValue] = selfNode

					// add the parent relationship
					for k := j+1; k<len(serializedString);k++ {
						if serializedString[k] == 'R' {
							// create the root node
							//fmt.Println("it's the root node, there is no parent relation")
							root = selfNode
							break
						}

						// reach the next node
						if serializedString[k] == 'N' {
							//fmt.Println(j+1,k)
							parentValue, _ := strconv.Atoi(string(serializedString[j+1:k]))
							//fmt.Println("find next node, the parentValue", parentValue)
							// add the parent relationship
							nodeMap[parentValue].childrens = append(nodeMap[parentValue].childrens, selfNode)
							break
						}

						// reach the end
						if k == len(serializedString)-1 {
							//fmt.Println(j+1,len(serializedString))
							parentValue, _ := strconv.Atoi(string(serializedString[j+1:len(serializedString)]))
							//fmt.Println("reach the end, parentValue", parentValue)
							// add the parent relationship
							nodeMap[parentValue].childrens = append(nodeMap[parentValue].childrens, selfNode)
							break
						}
					}

					break
				}
			}
		}
	} 

	return root
}

func main() {
	fmt.Println("Hello, playground")

	// create the test tree
	node5 := &NArrayNode{5, nil}
	node6 := &NArrayNode{6, nil}
	node8 := &NArrayNode{8, nil}
	node3 := &NArrayNode{3, []*NArrayNode{node5, node6}}
	node4 := &NArrayNode{2, []*NArrayNode{node8}}
	node2 := &NArrayNode{2, nil}
	node1 := &NArrayNode{1, []*NArrayNode{node3, node2, node4}}

	fmt.Println(Serialize(node1))
	root := DeSerialize(Serialize(node1))
	fmt.Println(Serialize(root))
}
