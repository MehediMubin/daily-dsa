// Problem: 3Sum
// Link: https://leetcode.com/problems/3sum/description/
// Time Complexity: O(n^2)
// Space Complexity: O(m) where m is the number of unique triplets that sum to 0

func threeSum(nums []int) [][]int {
  sort.Ints(nums)
  triplets := [][]int{}

  for i := 0; i + 2 < len(nums); i++ {
    if i > 0 && nums[i] == nums[i - 1] {
      continue
    }

    left, right := i + 1, len(nums) - 1
    for left < right {
      sum := nums[i] + nums[left] + nums[right]
      if sum == 0 {
        triplets = append(triplets, []int{nums[i], nums[left], nums[right]})

        for left < right && nums[left] == nums[left + 1] {
          left++
        }

        for left < right && nums[right] == nums[right - 1] {
          right--
        }

        left++
        right--
      } else if sum < 0 {
        left++
      } else {
        right--
      }
    }
  }

  return triplets
}