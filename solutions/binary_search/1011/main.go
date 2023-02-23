package main

func main() {
	println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, days int) int {
	total, max := 0, -1
	for _, weight := range weights {
		total += weight
		if weight > max {
			max = weight
		}
	}

	l, r, mid := max, total, 0
	for l < r {
		mid = (r + l) / 2
		d, s := 1, 0
		for _, weight := range weights {
			if s+weight > mid {
				d++
				s = 0
			}

			s += weight
		}

		if d > days {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
