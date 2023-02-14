package main

func main() {
	println(countOdds(4, 7))
}

func countOdds(low int, high int) int {
	res := (high - low + 1) / 2
	if low%2 == 1 && high%2 == 1 {
		res++
	}
	return res
}
