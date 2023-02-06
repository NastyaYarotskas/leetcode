package main

/*
Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].
*/

func main() {
	res := shuffle([]int{1, 1, 2, 2}, 2)
	for _, re := range res {
		println(re)
	}
}

func shuffle(nums []int, n int) []int {
	newArr := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		newArr = append(newArr, nums[i], nums[i+n])
	}
	return newArr
}
