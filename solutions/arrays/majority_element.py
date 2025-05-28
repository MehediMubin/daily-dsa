# Problem: Majority Element
# Link: https://leetcode.com/problems/majority-element/

# Time Complexity: O(n log n)
# Space Complexity: O(1)

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        nums.sort()
        return nums[len(nums) // 2]


# Boyer Moore's Voting Algorithm
# Time Complexity: O(n)
# Space Complexity: O(1)

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cnt = 0
        candidate = None

        for num in nums:
            if cnt == 0:
                candidate = num
            
            cnt += (1 if num == candidate else -1)
        
        return candidate