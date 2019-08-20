package main

import (
	"fmt"
)

func BFS(relation map[string][]string, startName string, searchName string) bool {
	alreadySearch := make(map[string]bool)
	prevSearch := make(map[string]string)
	prevSearch[startName] = ""
	searchQueue := make([]string, 0)
	searchQueue = append(searchQueue, startName)

	for len(searchQueue) != 0 {
		if _, ok := alreadySearch[searchQueue[0]]; !ok {
			if searchQueue[0] != searchName {
				for _, name := range relation[searchQueue[0]] {
					prevSearch[name] = searchQueue[0]

					if name == searchName {
						for searchName != "" {
							fmt.Println("Search node path:", searchName)
							searchName = prevSearch[searchName]
						}
						return true
					}
				}
				searchQueue = append(searchQueue, relation[searchQueue[0]]...)
				alreadySearch[searchQueue[0]] = true
			}
		}

		searchQueue = searchQueue[1:]
	}

	return false
}

func DFSRecursion(relation map[string][]string, startName string, searchName string) bool {
	if startName != searchName {
		for _, name := range relation[startName] {
			find := DFSRecursion(relation, name, searchName)
			fmt.Println("Search node path:", startName)
			if find {
				return true
			}
		}
	} else {
		fmt.Println("Search node path:", searchName)
		return true
	}
	return false
}

func main() {
	fmt.Println("Welcome to the playground!")
	releationMap := make(map[string][]string)
	releationMap["name A"] = []string{"name B", "name C"}
	releationMap["name B"] = []string{"name E", "name F"}
	releationMap["name C"] = []string{"name F", "name OK"}
	releationMap["name E"] = []string{"name F", "name OK"}
	releationMap["name F"] = []string{"name OK"}
	releationMap["name OK"] = nil

	fmt.Println("BFS search name OK start from name A:", BFS(releationMap, "name A", "name OK"))
	fmt.Println("DFS search name OK start from name A:", DFSRecursion(releationMap, "name A", "name OK"))
}
