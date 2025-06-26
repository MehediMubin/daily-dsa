// Problem: Contains Duplicate 2
// Link: https://leetcode.com/problems/contains-duplicate-ii/description/
// Time Complexity: O(n)
// Space Complexity: O(n)

func containsNearbyDuplicate(nums []int, k int) bool {
  mp := map[int]int{}
  for i, num := range nums {
    idx, ok := mp[num]
    if ok {
      diff := i - idx
      if diff <= k {
        return true
      }
    }
    mp[num] = i
  }   
  return false
}