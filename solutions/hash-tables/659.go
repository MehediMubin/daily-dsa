// Problem: Split Array into Consecutive Subsequences
// Link: https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Time Complexity: O(n)
// Space Complexity: O(n)

func isPossible(nums []int) bool {
  freq, subs := make(map[int]int), make(map[int]int)
  for _, num := range nums {
    freq[num]++
  }

  for _, num := range nums {
    if freq[num] == 0 {
      continue
    }

    if subs[num - 1] > 0 {
      subs[num - 1]--
      subs[num]++
    } else if freq[num + 1] > 0 && freq[num + 2] > 0 {
      freq[num + 1]--
      freq[num + 2]--
      subs[num + 2]++
    } else {
      return false
    }

    freq[num]--
  }

  return true
}