package main

/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search
target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	i := (r - l) / 2

	for l != r {
		currVal := nums[i]
		if currVal == target {
			return i
		} else if currVal > target {
			r = i
			i = (r - l) / 2
		} else {
			l = i + 1
			i = l + (r-l)/2
		}
	}

	if nums[i] == target {
		return i
	}

	return -1
}
