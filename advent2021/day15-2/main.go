package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inf = 10000000

type Item struct {
	index    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(v interface{}) {
	queue := *pq
	n := len(queue)
	queue = queue[:n+1]
	item := v.(*Item)
	queue[n] = item
	*pq = queue
}

func (pq *PriorityQueue) Pop() interface{} {
	queue := *pq
	n := len(queue)
	item := queue[n-1]
	*pq = queue[:n-1]
	return item
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var risks [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, v := range scanner.Text() {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalf("not and int: %v", err)
			}
			row = append(row, i)
		}
		risks = append(risks, row)
	}

	newRisks := expandRisks(risks)

	n, path := dijkstra(newRisks)

	var strMatrix [][]string
	for _, r := range newRisks {
		var row []string
		for _, v := range r {
			row = append(row, strconv.Itoa(v)+" ")
		}
		strMatrix = append(strMatrix, row)
	}

	for _, v := range path {
		strMatrix[v/n][v%n] = strings.Replace(strMatrix[v/n][v%n], " ", "!", -1)
	}

	for _, v := range strMatrix {
		fmt.Println(v)
	}
}

func dijkstra(risks [][]int) (int, []int) {
	m := len(risks)
	n := len(risks[0])

	graph := make(map[int][][]int, 0)

	for i, row := range risks {
		for j := range row {
			if i-1 >= 0 {
				graph[i*n+j] = append(graph[i*n+j], []int{(i-1)*n + j, risks[i-1][j]})
			}
			if i+1 <= m-1 {
				graph[i*n+j] = append(graph[i*n+j], []int{(i+1)*n + j, risks[i+1][j]})
			}
			if j-1 >= 0 {
				graph[i*n+j] = append(graph[i*n+j], []int{i*n + j - 1, risks[i][j-1]})
			}
			if j+1 <= n-1 {
				graph[i*n+j] = append(graph[i*n+j], []int{i*n + j + 1, risks[i][j+1]})
			}
		}
	}

	dist := make([]int, m*n)

	dist[0] = 0
	for i := 1; i < len(dist); i++ {
		dist[i] = inf
	}

	parents := make([]int, m*n)
	for i := 1; i < len(parents); i++ {
		parents[i] = -1
	}

	pq := make(PriorityQueue, 0, n*m)
	started := &Item{
		index:    0,
		priority: 0,
	}
	heap.Push(&pq, started)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.priority > dist[item.index] {
			continue
		}

		for _, slice := range graph[item.index] {
			idx, val := slice[0], slice[1]
			distance := val + dist[item.index]
			if distance < dist[idx] {
				dist[idx] = distance
				heap.Push(&pq, &Item{index: idx, priority: distance})
				parents[idx] = item.index
			}
		}
	}

	fmt.Println("the shortest way =", dist[len(dist)-1])

	path := make([]int, 0)
	for v := len(parents) - 1; v != 0; v = parents[v] {
		path = append(path, v)
	}

	path = append(path, 0)

	return n, path
}

func expandRisks(risks [][]int) [][]int {
	m, n := len(risks), len(risks[0])

	for i, row := range risks {
		for block := 1; block < 5; block++ {
			for _, v := range row[:n] {
				risks[i] = append(risks[i], ((v+block-1)%9)+1)
			}
		}
	}

	for block := 1; block < 5; block++ {
		for _, row := range risks[:m] {
			temp := make([]int, 0)
			for _, v := range row {
				temp = append(temp, ((v+block-1)%9)+1)
			}
			risks = append(risks, temp)
		}
	}

	return risks
}
