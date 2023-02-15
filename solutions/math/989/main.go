package main

func main() {
	res := addToArrayForm([]int{2, 1, 5}, 806)
	for _, re := range res {
		print(re)
	}
}

func addToArrayForm(num []int, k int) []int {
	res := make([]int, 0, len(num))
	i, sum, carry := len(num)-1, 0, 0
	for k > 0 || i >= 0 || carry > 0 {
		sum = carry
		if i >= 0 {
			sum += num[i]
			i--
		}
		if k != 0 {
			sum += k % 10
			k /= 10
		}
		res = append(res, sum%10)
		carry = sum / 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
