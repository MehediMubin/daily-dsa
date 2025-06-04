// Problem: Rotate Array
// Link: https://leetcode.com/problems/rotate-array/
// Time Complexity: O(n)
// Space Complexity: O(n)

func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    pos := n - k
    if k == 0 {
        return
    }
    var arr []int
    for i := pos; i < n; i++ {
        arr = append(arr, nums[i])
    }

    for i := 0; i < pos; i++ {
        arr = append(arr, nums[i])
    }
    copy(nums, arr)
}