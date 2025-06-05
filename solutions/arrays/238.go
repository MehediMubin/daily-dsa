// Problem: Product of Array Except Self
// Link: https://leetcode.com/problems/product-of-array-except-self/
// Time Complexity: O(n)
// Space Complexity: O(n)

func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    suffix := make([]int, n)

    prefix[0] = nums[0]
    suffix[n - 1] = nums[n - 1]
    for i, j := 1, n - 2; i + 1 < n; i, j = i + 1, j - 1 {
        prefix[i] = prefix[i - 1] * nums[i]
        suffix[j] = suffix[j + 1] * nums[j]
    }

    nums[0] = suffix[1]
    nums[n - 1] = prefix[n - 2]
    for i := 1; i + 1 < n; i++ {
        nums[i] = prefix[i - 1] * suffix[i + 1]
    }
    return nums
}