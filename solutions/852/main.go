package main

/*
An array arr a mountain if the following properties hold:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given a mountain array arr, return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

You must solve it in O(log(arr.length)) time complexity.
*/

func main() {
	println(peakIndexInMountainArray([]int{0, 2, 3, 4, 5, 10, 5, 2}))
}

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1
	mid := (r - l) / 2

	for l != r {
		prevVal := arr[mid-1]
		curVal := arr[mid]
		nextVal := arr[mid+1]
		if prevVal < curVal && curVal > nextVal {
			return mid
		} else if prevVal < curVal && curVal < nextVal {
			l = mid
			mid = l + (r-l)/2
		} else if prevVal > curVal && curVal > nextVal {
			r = mid
			mid = (r - l) / 2
		}
	}

	return mid
}
