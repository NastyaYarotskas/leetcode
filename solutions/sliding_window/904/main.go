package main

func main() {
	println(totalFruit([]int{1, 0, 1, 4, 1, 4, 1, 2, 3}))
}

func totalFruit(fruits []int) int {
	lastIndex, val := 0, 0
	maxSum := 1
	sum := 1
	m := make(map[int][]int)
	m[fruits[0]] = append(m[fruits[0]], 0)

	for i := 1; i < len(fruits); i++ {
		if _, ok := m[fruits[i]]; ok {
			m[fruits[i]] = append(m[fruits[i]], i)
			sum++
		} else if _, ok := m[fruits[i]]; !ok && len(m) == 1 {
			m[fruits[i]] = append(m[fruits[i]], i)
			sum++
		} else {
			prev := fruits[i-1]
			for key, values := range m {
				if key != prev && key != fruits[i] {
					if sum > maxSum {
						maxSum = sum
					}

					val = key
					sum -= len(values)
					lastIndex = values[len(values)-1]
					break
				}
			}
			delete(m, val)
			prevInx := make([]int, 0)
			for _, elem := range m[prev] {
				if elem > lastIndex {
					prevInx = append(prevInx, elem)
				} else {
					sum--
				}
			}
			m[prev] = prevInx
			m[fruits[i]] = append(m[fruits[i]], i)
			sum++
		}
	}

	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}
