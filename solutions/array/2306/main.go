package main

func main() {
	println(distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
}

func distinctNames(ideas []string) int64 {
	firstLetterToPrefix := make(map[rune]map[string]struct{}, len(ideas))
	for _, idea := range ideas {
		prefix := idea[1:]
		firstLetter := []rune(idea)[0]

		if _, ok := firstLetterToPrefix[firstLetter]; !ok {
			firstLetterToPrefix[firstLetter] = make(map[string]struct{})
		}

		firstLetterToPrefix[firstLetter][prefix] = struct{}{}
	}

	var count, same = 0, 0
	for l1, set1 := range firstLetterToPrefix {
		for l2, set2 := range firstLetterToPrefix {
			if l1 >= l2 {
				continue
			}
			same = 0
			for l := range set2 {
				if _, ok := set1[l]; ok {
					same++
				}
			}

			count += (len(set1) - same) * (len(set2) - same)
		}
	}

	return int64(count) * 2
}
