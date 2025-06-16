// Problem: Single Number
// Link: https://leetcode.com/problems/single-number/description/
// Time Complexity: O(n)
// Space Complexity: O(1)

func singleNumber(nums []int) int {
    ans := 0
    for _, num := range nums {
      ans ^= num
    }
    return ans
}