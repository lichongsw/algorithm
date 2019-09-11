package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority, so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func GetMin(first int, second int) int {
	result := first
	if result > second {
		result = second
	}

	return result
}

type Node struct {
	name        string
	weight      int
}

func Dijkstra(relation map[string][]Node, startName string, searchName string) (int, []string) {
	searchQueue := make(PriorityQueue, 1)
	index := 0
	searchQueue[index] = &Item{startName, 0, index}
	index++
	heap.Init(&searchQueue)

	preNodeMap := make(map[string]string)
	searchedMap := make(map[string]*Item)
	for len(searchQueue) != 0 {
		currentNode := heap.Pop(&searchQueue).(*Item)
		currentName := currentNode.value
		if currentName == searchName {
			break
		}

		//fmt.Println("Debug current node", currentName)
		for _, node := range relation[currentName] {
			if item, ok := searchedMap[node.name]; ok {
				if item.priority > currentNode.priority+node.weight {
					preNodeMap[node.name] = currentName
				}
				searchQueue.Update(item, node.name, GetMin(item.priority, currentNode.priority+node.weight))
				//fmt.Println("Debug update exist node:", node.name, item)
			} else {
				newItem := &Item{node.name, node.weight + currentNode.priority, index}
				index++
				searchedMap[node.name] = newItem
				//fmt.Println("Debug add new node:", node.name, newItem, &newItem)
				heap.Push(&searchQueue, newItem)
				preNodeMap[node.name] = currentName
			}
		}
	}
	
	path := []string{}
	pathName := searchName
	for pathName != startName {
		if name, ok := preNodeMap[pathName]; ok {
			path = append([]string{pathName}, path...)
			pathName = name
		}
	}
	path = append([]string{pathName}, path...)

	return searchedMap[searchName].priority, path
}

func main() {
	fmt.Println("Welcome to the playground!")
	releationMap := make(map[string][]Node)
	releationMap["name A"] = []Node{Node{"name B", 5}, Node{"name E", 20}}
	releationMap["name B"] = []Node{Node{"name C", 10}}
	releationMap["name C"] = []Node{Node{"name D", 20}, Node{"name E", 4}}
	releationMap["name D"] = []Node{Node{"name F", 5}}
	releationMap["name E"] = []Node{Node{"name D", 10}}
	releationMap["name F"] = nil

	distance, path := Dijkstra(releationMap, "name A", "name D")
	fmt.Println("Dijkstra search name D start from name A:", distance, "with path:", path)
	
	releationMap["name C"] = []Node{Node{"name D", 20}, Node{"name E", 6}}
	distance, path = Dijkstra(releationMap, "name A", "name D")
	fmt.Println("Dijkstra search name D start from name A:", distance, "with path:", path)
}

// Welcome to the playground!
// Dijkstra search name D start from name A: 29 with path: [name A name B name C name E name D]
// Dijkstra search name D start from name A: 30 with path: [name A name E name D]
