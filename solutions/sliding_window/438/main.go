package main

/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in
any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the
original letters exactly once.
*/

func main() {
	res := findAnagrams("cbaebabacd", "abc")
	for _, re := range res {
		println(re)
	}
}

func findAnagrams(s string, p string) []int {
	startIndexes := make([]int, 0)
	pLen := len(p)
	if len(s) < len(p) {
		return startIndexes
	}

	pFreq := make(map[string]int, len(p))
	for _, val := range p {
		pFreq[string(val)] += 1
	}

	sFreq := make(map[string]int, len(s[:pLen]))
	for _, val := range s[:pLen] {
		sFreq[string(val)] += 1
	}

	for i := 0; i < len(s)-pLen+1; i++ {
		f := true
		for k := range pFreq {
			if sFreq[k] != pFreq[k] {
				f = false
				break
			}
		}
		if f {
			startIndexes = append(startIndexes, i)
		}
		sFreq[string(s[i])]--
		if i+pLen < len(s) {
			sFreq[string(s[i+pLen])]++
		}
	}

	return startIndexes
}
