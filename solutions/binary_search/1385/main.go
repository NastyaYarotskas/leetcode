package main

import (
	"sort"
)

func main() {
	res := findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)
	println(res)
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	res := 0
	for _, i := range arr1 {
		if !checkDistance(i, arr2, d) {
			res++
		}
	}

	return res
}

func checkDistance(v int, arr2 []int, d int) bool {
	l, r := 0, len(arr2)-1
	for l <= r {
		mid := l + (r-l)/2
		if (arr2[mid]-v)*(arr2[mid]-v) <= d*d {
			return true
		} else if arr2[mid] < v {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
