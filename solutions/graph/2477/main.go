package main

import "math"

func main() {
	res := minimumFuelCost([][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}}, 2)
	//res := minimumFuelCost([][]int{{0, 1}, {0, 2}, {0, 3}}, 5)
	//res := minimumFuelCost([][]int{}, 1)
	println(res)
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	graph := make(map[int][]int)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}

	used := make(map[int]bool)

	_, f := dfs(0, used, graph, seats)

	return int64(f)
}

func dfs(v int, used map[int]bool, graph map[int][]int, seats int) (int, int) {
	used[v] = true
	people := 1
	fuel := 0
	for _, u := range graph[v] {
		if !used[u] {
			p, f := dfs(u, used, graph, seats)
			people += p
			fuel += f
		}

	}

	if v != 0 {
		fuel += int(math.Ceil(float64(people) / float64(seats)))
	}

	return people, fuel
}
