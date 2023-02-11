package main

func main() {
	for _, i := range shortestAlternatingPaths(5,
		[][]int{{2, 2}, {0, 1}, {0, 3}, {0, 0}, {0, 4}, {2, 1}, {2, 0}, {1, 4}, {3, 4}},
		[][]int{{1, 3}, {0, 0}, {0, 3}, {4, 2}, {1, 0}}) {
		println(i)
	}
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	defColor, redColor, blueColor := 0, 1, 2
	type vertex struct {
		v         int
		prevColor int
	}

	graph := make(map[int][]*vertex, n)
	for i := 0; i < n; i++ {
		graph[i] = []*vertex{}
	}
	for _, edge := range redEdges {
		graph[edge[0]] = append(graph[edge[0]], &vertex{prevColor: redColor, v: edge[1]})
	}
	for _, edge := range blueEdges {
		graph[edge[0]] = append(graph[edge[0]], &vertex{prevColor: blueColor, v: edge[1]})
	}

	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, -1)
	}
	ans[0] = 0

	q := make([]vertex, 0, n)
	qHead, qTail := 0, 0
	q = append(q, vertex{prevColor: defColor, v: 0})
	step := 1
	for qHead <= qTail {
		dif := qTail - qHead
		for i := 0; i <= dif; i++ {
			v := q[qHead]
			qHead++

			for _, to := range graph[v.v] {
				if to.v != -1 && v.prevColor != to.prevColor {
					if ans[to.v] == -1 {
						ans[to.v] = step
					}
					qTail++
					q = append(q, vertex{v: to.v, prevColor: to.prevColor})
					to.v = -1
				}
			}
		}
		step++
	}

	return ans
}
