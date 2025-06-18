// Problem: Counting Bits
// Link: https://leetcode.com/problems/counting-bits/description/
// Time Complexity: O(n log n)
// Space Complexity: O(n)

func solve(n int) int {
  cnt := 0
  for n > 0 {
    if n % 2 != 0 {
      cnt++
    }
    n /= 2
  }
  return cnt
}

func countBits(n int) []int {
    ans := []int{}
    for i := 0; i <= n; i++ {
      ans = append(ans, solve(i))
    }
    return ans
}