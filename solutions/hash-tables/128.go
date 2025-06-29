// Problem: Longest Consecutive Sequence
// Link: https://leetcode.com/problems/longest-consecutive-sequence/description/
// Time Complexity: O(n log n)
// Space Complexity: O(1)

func longestConsecutive(nums []int) int {
    sort.Ints(nums)
    ans, counter, expectedValue := 0, 0, 0
    for i, num := range nums {
      if i == 0 {
        counter++
        expectedValue = num + 1
        continue
      }
      if num == expectedValue {
        counter++
        expectedValue++
      } else if num == expectedValue - 1 {
        continue
      } else {
        ans = max(ans, counter)
        counter = 1
        expectedValue = num + 1
      }
    }
    ans = max(ans, counter)

    return ans
}