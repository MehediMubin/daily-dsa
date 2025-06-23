// Problem: Number of Good Pairs 
// Link: https://leetcode.com/problems/number-of-good-pairs/description/
// Time Complexity: O(n)
// Space Complexity: O(n)

func numIdenticalPairs(nums []int) int {
  mp := map[int]int{}
  for _, num := range nums {
    mp[num]++
  }

  ans := 0
  for _, value := range mp {
    value--
    ans += value * (value + 1) / 2
  }
  return ans
}