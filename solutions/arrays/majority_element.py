# Problem: Majority Element
# Link: https://leetcode.com/problems/majority-element/

# Time Complexity: O(n log n)
# Space Complexity: O(1)

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        nums.sort()
        return nums[len(nums) // 2]