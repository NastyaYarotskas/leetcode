package main

func main() {
	println(maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
	println(maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
	println(maxDistance([][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	queue := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if len(queue) == n*n || len(queue) == 0 {
		return -1
	}

	distance := 0
	l, r := 0, len(queue)

	for {
		for _, point := range queue[l:r] {
			i := point[0]
			j := point[1]
			if j+1 < n && grid[i][j+1] == 0 {
				grid[i][j+1] = 1
				queue = append(queue, []int{i, j + 1})
			}
			if j-1 >= 0 && grid[i][j-1] == 0 {
				grid[i][j-1] = 1
				queue = append(queue, []int{i, j - 1})
			}
			if i+1 < n && grid[i+1][j] == 0 {
				grid[i+1][j] = 1
				queue = append(queue, []int{i + 1, j})
			}
			if i-1 >= 0 && grid[i-1][j] == 0 {
				grid[i-1][j] = 1
				queue = append(queue, []int{i - 1, j})
			}
		}
		if len(queue) == r {
			break
		}
		l = r
		r = len(queue)
		distance++
	}

	return distance
}
