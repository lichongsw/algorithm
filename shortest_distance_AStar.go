package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string  // The value of the item; arbitrary.
	priority float64 // The priority of the item in the queue.
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
func (pq *PriorityQueue) Update(item *Item, value string, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func GetFloat64Min(first float64, second float64) float64 {
	result := first
	if result > second {
		result = second
	}

	return result
}

type Node struct {
	name   string
	weight float64
}

type Vetex2 struct {
	x int
	y int
}

func ManhattanDistance(a, b Vetex2) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
}

func AStar(startName string, searchName string, relation map[string][]Node, Vetexs map[string]Vetex2) (float64, []string) {
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
		fmt.Println("Debug current node", currentName)
		if currentName == searchName {
			break
		}

		for _, node := range relation[currentName] {
			currentWeight := currentNode.priority + node.weight
			if item, ok := searchedMap[node.name]; ok {
				if item.priority + ManhattanDistance(Vetexs[node.name], Vetexs[searchName]) > currentWeight + ManhattanDistance(Vetexs[currentName], Vetexs[searchName]) {
					preNodeMap[node.name] = currentName
				}
				searchQueue.Update(item, node.name, GetFloat64Min(item.priority, currentWeight))
				//fmt.Println("Debug update exist node:", node.name, item)
			} else {
				newItem := &Item{node.name, currentWeight, index}
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
	releationMap["name A"] = []Node{Node{"name B", 8}, Node{"name D", 10}}
	releationMap["name B"] = []Node{Node{"name C", 3}}
	releationMap["name C"] = []Node{Node{"name D", 3}, Node{"name E", 4}}
	releationMap["name D"] = []Node{Node{"name E", 5}}
	releationMap["name E"] = nil

	vetex2Map := make(map[string]Vetex2)
	vetex2Map["name A"] = Vetex2{0, 0}
	vetex2Map["name B"] = Vetex2{8, 0}
	vetex2Map["name C"] = Vetex2{8, 3}
	vetex2Map["name D"] = Vetex2{8, 6}
	vetex2Map["name E"] = Vetex2{12, 3}

	distance, path := AStar("name A", "name E", releationMap, vetex2Map)
	fmt.Println("AStar search name E start from name A:", distance, "with path:", path)
}

// Welcome to the playground!
// Debug current node name A
// Debug current node name B
// Debug current node name D
// Debug current node name C
// Debug current node name E
// AStar search name D start from name A: 15 with path: [name A name D name E]