package main

func main() {
	println(isPerfectSquare(4))
}

func isPerfectSquare(num int) bool {
	l, r := 0, num

	for l <= r {
		mid := l + (r-l)/2
		sqrt := mid * mid
		if sqrt == num {
			return true
		}
		if sqrt < num {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return false
}
