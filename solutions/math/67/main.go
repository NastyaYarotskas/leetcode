package main

import "strconv"

func main() {
	println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	a = reverse(a)
	b = reverse(b)
	c := ""
	carry := 0
	sum := 0
	i, j := 0, 0

	for i < len(a) || j < len(b) || carry != 0 {
		sum = carry
		if i < len(a) {
			v, _ := strconv.Atoi(string(a[i]))
			sum += v
			i++
		}
		if j < len(b) {
			v, _ := strconv.Atoi(string(b[j]))
			sum += v
			j++
		}
		c = strconv.Itoa(sum%2) + c
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
	}

	return c
}

func reverse(s string) string {
	res := ""
	for _, el := range s {
		res = string(el) + res
	}
	return res
}
