package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	index int
	name int
	sta_list []int //the list of this node start to others node
	cost_list []int //the list of cost of with the connect node
	distance int
	PI int
}
type Edge struct {
	start int
	end  int
	weight int
}
//**************************** PriorityQueue
//
//
//***************************
type PriorityQueue []*Node
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1
	*pq = old[0 : n-1]
	return node
}
func (pq *PriorityQueue) update(node *Node, distance int) {
	node.distance = distance
	heap.Fix(pq, node.index)
}
//////////////////////////////////////////////
var nodes_sum int
var edges_sum int
var source int
var set_in []int
func main()  {
	var e Edge
	fmt.Scanf("%d %d %d", &nodes_sum, &edges_sum, &source)
	nodes := make([]Node, nodes_sum, 100000)
	edges :=make([]Edge,edges_sum,1000000)
	for i := 0; i < edges_sum; i++ {
		fmt.Scanf("%d %d %d",&e.start, &e.end,&e.weight)
		nodes[e.start].sta_list = append(nodes[e.start].sta_list, e.end)
		nodes[e.start].cost_list = append(nodes[e.start].cost_list, e.weight)
	}
	for i := 0; i < nodes_sum; i++ {
		nodes[i].name=i
		nodes[i].index=i
		nodes[i].distance=100000
		nodes[i].PI=-1
		if i==source {
			nodes[i].distance=0
		}
	}
	Dijkstra(nodes,edges)
	for i := 0; i < nodes_sum; i++ {
		if nodes[i].distance == 100000 {
			fmt.Println("INF")
		}else {
			fmt.Println(nodes[i].distance)
		}

	}
}

func Dijkstra(nodes []Node,edges []Edge) {
	pq := make(PriorityQueue, nodes_sum)
	for i:=0; i< len(nodes);i++{
		pq[i] = & Node{
			sta_list: nodes[i].sta_list,
			cost_list:nodes[i].cost_list,
			distance: nodes[i].distance,
			name:nodes[i].name,
			PI: nodes[i].PI,
			index: nodes[i].index,
		}
	}
	heap.Init(&pq)
	for ;len(pq)!=0 ;  {
		n := heap.Pop(&pq).(*Node)
		for i := 0; i< len(n.sta_list);i++  {
			temp := nodes[n.sta_list[i]]
			weight := RELAX(*n,temp,nodes)
			temp.distance=weight
			if len(pq)==0 {
				return
			}
			temp.index=findIndex(pq,temp)
			if weight != 1000000 && temp.distance!=0{
				pq[temp.index].distance=temp.distance
				pq.update(&temp,temp.distance)
			}
		}
	}
}
func RELAX(start Node,end Node,nodes []Node) int {
	weight := GetWeight(start,end,nodes)
	if end.distance>start.distance+weight {
		nodes[end.name].distance=start.distance+weight
		nodes[end.name].PI=start.name
	}
	return nodes[end.name].distance
}
func GetWeight(start Node,end Node,nodes []Node) int {
	for i := 0; i<len(start.sta_list);i++  {
		if end.name == start.sta_list[i] {
			return start.cost_list[i]
		}
	}
	return 100000
}
func findIndex(pq PriorityQueue,temp Node) int{
	for i:=0; i<len(pq) ;i++ {
		if temp.name == pq[i].name {
			return pq[i].index
		}
	}
	return 0
}