// Problem: Number of Zero-Filled Subarrays
// Link: https://leetcode.com/problems/number-of-zero-filled-subarrays/
// Time Complexity: O(n)
// Space Complexity: O(1)

func zeroFilledSubarray(nums []int) int64 {
    var ans int64 = 0
    cnt := 0
    for _, num := range nums {
        if num == 0 {
            cnt++
        } else {
            ans += int64(cnt * (cnt + 1)) / 2
            cnt = 0
        }
    }

    ans += int64(cnt * (cnt + 1)) / 2
    return ans
}