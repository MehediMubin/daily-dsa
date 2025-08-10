// Problem: Two Sum
// Link: https://leetcode.com/problems/two-sum/
// Time Complexity: O(n log n)
// Space Complexity: O(n) 

func twoSum(nums []int, target int) []int {
  mp := make(map[int][]int)
  for i, num := range nums {
    mp[num] = append(mp[num], i)
  }
  
  sort.Ints(nums)
  l, r := 0, len(nums) - 1
  for l < r {
    sum := nums[l] + nums[r]
    if sum == target {
      if nums[l] == nums[r] {
        return []int{mp[nums[l]][0], mp[nums[l]][1]}
      } else {
        return []int{mp[nums[l]][0], mp[nums[r]][0]}
      }
    } else if sum < target {
      l++
    } else {
      r--
    }
  }
  return []int{}
}