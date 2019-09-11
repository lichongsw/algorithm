package main

import (
	"fmt"
)

type Node struct {
	preNodeName string
	name        string
	weight      int
}

func BellmanFord(relation map[string][]Node, startName string, searchName string) (int, []string) {
	searchQueue := make([]string, 0)
	searchQueue = append(searchQueue, startName)
	searchedMap := make(map[string]Node)

	for len(searchQueue) != 0 {
		currentName := searchQueue[0]
		//fmt.Println("Debug current node", currentName)
		for _, node := range relation[currentName] {
			if existNode, ok := searchedMap[node.name]; ok {
				if searchedMap[node.name].weight > node.weight+searchedMap[currentName].weight {
					existNode.weight = node.weight + searchedMap[currentName].weight
					existNode.preNodeName = currentName
					searchedMap[node.name] = existNode
					//fmt.Println("Debug update exist node:", node.name, searchedMap[node.name])
				}
			} else {
				searchedMap[node.name] = Node{currentName, node.name, node.weight + searchedMap[currentName].weight}
				fmt.Println("Debug add new node:", node.name, searchedMap[node.name])
				searchQueue = append(searchQueue, node.name)
			}
		}

		searchQueue = searchQueue[1:]
	}

	fmt.Println("Debug searched record:", searchedMap)
	path := []string{}
	pathName := searchName
	for pathName != startName {
		if pathNode, ok := searchedMap[pathName]; ok {
			path = append([]string{pathName}, path...)
			pathName = pathNode.preNodeName
		}
	}
	path = append([]string{pathName}, path...)

	return searchedMap[searchName].weight, path
}

func main() {
	fmt.Println("Welcome to the playground!")
	releationMap := make(map[string][]Node)
	releationMap["name A"] = []Node{Node{"name A", "name B", 5}, Node{"name A", "name E", 0}}
	releationMap["name B"] = []Node{Node{"name B", "name C", 15}, Node{"name B", "name F", 20}}
	releationMap["name C"] = []Node{Node{"name C", "name D", 20}}
	releationMap["name D"] = nil
	releationMap["name E"] = []Node{Node{"name E", "name C", 30}, Node{"name E", "name F", 35}}
	releationMap["name F"] = []Node{Node{"name F", "name D", 10}}

	distance, path := BellmanFord(releationMap, "name A", "name D")
	fmt.Println("BellmanFord search name D start from name A:", distance, "with path:", path)
}

// Welcome to the playground!
// Debug add new node: name B {name A name B 5}
// Debug add new node: name E {name A name E 0}
// Debug add new node: name C {name B name C 20}
// Debug add new node: name F {name B name F 25}
// Debug add new node: name D {name C name D 40}
// Debug searched record: map[name B:{name A name B 5} name C:{name B name C 20} name D:{name F name D 35} name E:{name A name E 0} name F:{name B name F 25}]
// BellmanFord search name D start from name A: 35 with path: [name A name B name F name D]