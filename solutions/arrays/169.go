// Problem: Majority Element
// Link: https://leetcode.com/problems/majority-element/

// Time Complexity: O(n log n)
// Space Complexity: O(1)

import "sort"

func majorityElement(nums []int) int {
    sort.Ints(nums)
    idx := int(len(nums) / 2)
    return nums[idx]
}


// Boyer Moore's Voting Algorithm
// Time Complexity: O(n)
// Space Complexity: O(1)

func majorityElement(nums []int) int {
    cnt := 0
    candidate := 0
    for _, num := range(nums) {
        if cnt == 0 {
            candidate = num
        } 
        if (candidate == num) {
            cnt++
        } else {
            cnt--
        }
    }
    return candidate
}