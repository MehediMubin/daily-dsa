// Problem: Number of Good Ways to Split a String
// Link: https://leetcode.com/problems/number-of-good-ways-to-split-a-string/description/
// Time Complexity: O(n)
// Space Complexity: O(n)

func numSplits(s string) int {
	freqLeft := make(map[int]int)
	freqRight := make(map[int]int)
	freq := make(map[byte]bool)

	distinct := 0
	for i := 0; i + 1 < len(s); i++ {
		if freq[s[i]] == false {
			distinct++
		}
		freqLeft[i] = distinct
		freq[s[i]] = true
	}

	freq = make(map[byte]bool)
	distinct = 0
	for i := len(s) - 1; i > 0; i-- {
		if freq[s[i]] == false {
			distinct++
		}
		freqRight[i] = distinct
		freq[s[i]] = true
	}

	count := 0
	for i := 0; i + 1 < len(s); i++ {
		if freqLeft[i] == freqRight[i + 1] {
			count++
		}
	}

	return count
}