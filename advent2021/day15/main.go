package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inf = 10000000

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	risks := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []int{}
		for _, v := range scanner.Text() {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalf("not and int: %v", err)
			}
			row = append(row, i)
		}
		risks = append(risks, row)
	}
	m := len(risks)
	n := len(risks[0])

	graph := [][]int{}
	for i := 0; i < m*n; i++ {
		graph = append(graph, make([]int, m*n))
	}

	for i, row := range risks {
		for j, v := range row {
			if i-1 >= 0 {
				graph[(i-1)*n+j][i*n+j] = v
			}
			if i+1 <= m-1 {
				graph[(i+1)*n+j][i*n+j] = v
			}
			if j-1 >= 0 {
				graph[i*n+j-1][i*n+j] = v
			}
			if j+1 <= n-1 {
				graph[i*n+j+1][i*n+j] = v
			}
		}
	}

	dist := make([]int, m*n)
	visited := make([]bool, m*n)

	dist[0] = 0
	for i := 1; i < len(dist); i++ {
		dist[i] = inf
	}

	for {
		minIdx := -1
		minDist := inf
		for i, v := range dist {
			if v < minDist && !visited[i] {
				minDist = v
				minIdx = i
			}
		}

		if minIdx == -1 {
			break
		}

		for i, v := range graph[minIdx] {
			if v > 0 && v+dist[minIdx] < dist[i] {
				dist[i] = v + dist[minIdx]
			}
		}
		visited[minIdx] = true
	}
	fmt.Println(dist[len(dist)-1])
}
