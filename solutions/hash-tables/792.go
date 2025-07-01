// Problem: Number of Matching Subsequences
// Link: https://leetcode.com/problems/number-of-matching-subsequences/description/

// Time Complexity: O(L + n * m * log L)
// Space Complexity: O(L)
// where, L = length of s, n = length of words, and m = length of word in words

func numMatchingSubseq(s string, words []string) int {
	charIndices := make(map[byte][]int)
  for i := 0; i < len(s); i++ {
    charIndices[s[i]] = append(charIndices[s[i]], i)
  }

  count := 0
  for _, word := range words {
    prevIndex := -1
    matched := true

    for i := 0; i < len(word); i++ {
      ch := word[i]
      indices, ok := charIndices[ch]
      if !ok {
        matched = false
        break
      }

      pos := sort.SearchInts(indices, prevIndex + 1)
      if pos == len(indices) {
        matched = false
        break
      }

      prevIndex = indices[pos]
    }
    if matched {
      count++
    }
  }

  return count
}