package main

/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

*/

func main() {
	println(guessNumber(1))
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	if n == 1 {
		return n
	}

	l := 1
	r := n
	i := (r - l) / 2

	for l != r {
		attempt := guess(i)
		if attempt == 0 {
			return i
		} else if attempt == -1 {
			r = i
			i = (r - l) / 2
		} else {
			l = i + 1
			i = l + (r-l)/2
		}
	}

	return i
}
