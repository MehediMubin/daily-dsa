// Problem: Number of 1 Bits
// Link: https://leetcode.com/problems/number-of-1-bits/
// Time Complexity: O(log2(n))
// Space Complexity: O(1)

func hammingWeight(n int) int {
    ans := 0
    for n > 0 {
      if n % 2 != 0 {
        ans++
      }
      n /= 2
    }
    return ans
}