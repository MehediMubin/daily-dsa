// Problem: Two Sum 2 - Input Array is Sorted
// Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
// Time Complexity: O(n)
// Space Complexity: O(1)

func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    for l < r {
      sum := numbers[l] + numbers[r]
      if sum < target {
        l++
      } else if sum > target {
        r--
      } else {
        return []int{l + 1, r + 1}
      }
    }
    return []int{} // unreachable because the problem guarantees one solution
}