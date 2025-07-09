// Problem: Subarray Sum Equals K
// Link: https://leetcode.com/problems/subarray-sum-equals-k/
// Time Complexity: O(n^2)
// Space Complexity: O(1)

func subarraySum(nums []int, k int) int {
	var count, sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}