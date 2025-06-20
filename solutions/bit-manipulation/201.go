// Problem: Bitwise AND of Numbers Range
// Link: https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
// Time Complexity: O(n - m)
// Space Complexity: O(1)

func rangeBitwiseAnd(left int, right int) int {
    ans := left
    for i := left + 1; i <= right; i++ {
      ans &= i
      if ans == 0 {
        break
      }
    }
    return ans
}