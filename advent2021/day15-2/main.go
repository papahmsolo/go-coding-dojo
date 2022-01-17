package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inf = 10000000

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

	var graph [][]int
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

	parents := make([]int, m*n)
	for i := 1; i < len(parents); i++ {
		parents[i] = -1
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
				parents[i] = minIdx
			}
		}

		visited[minIdx] = true
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
	m := len(risks)
	for i, row := range risks {
		for block := 1; block < 5; block++ {
			for _, v := range row[:m] {
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
