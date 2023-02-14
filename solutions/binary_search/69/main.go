package main

func main() {
	println(mySqrt(8))
}

func mySqrt(x int) int {
	l, r := 0, x
	res := 0
	for l <= r {
		mid := l + (r-l)/2
		res = mid
		if mid*mid <= x && x < (mid+1)*(mid+1) {
			res = mid
			break
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
