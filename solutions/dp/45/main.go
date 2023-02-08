package main

func main() {
	println(jump([]int{2, 3, 0, 1, 4}))
}

func jump(nums []int) int {
	minJumps := make([]int, len(nums), len(nums))
	fromJumps := make([][]int, len(nums), len(nums))

	for i, num := range nums {
		for j := 1; j <= num; j++ {
			if i+j < len(nums) {
				fromJumps[i+j] = append(fromJumps[i+j], i)
			}
		}
		currJump := make([]int, 0, len(fromJumps[i]))
		for _, el := range fromJumps[i] {
			currJump = append(currJump, minJumps[el])
		}
		minJumps[i] = findMinValue(currJump) + 1
	}

	return minJumps[len(nums)-1]
}

func findMinValue(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
