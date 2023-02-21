package main

func main() {
	println(singleNonDuplicate([]int{2}))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l, r, mid := 0, len(nums)-1, 0
	for l <= r {
		mid = l + (r-l)/2
		if mid == 0 && nums[mid] != nums[mid+1] {
			return nums[mid]
		}
		if mid == len(nums)-1 && nums[mid-1] != nums[mid] {
			return nums[mid]
		}
		if nums[mid-1] != nums[mid] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if nums[mid-1] == nums[mid] {
			if mid%2 == 1 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		if nums[mid+1] == nums[mid] {
			if mid%2 == 0 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return 0
}
