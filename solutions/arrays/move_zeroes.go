//  Problem: Move Zeroes
//  Link: https://leetcode.com/problems/move-zeroes/
//  Time: O(n)
//  Space: O(1)

func moveZeroes(nums []int)  {
    pos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[i], nums[pos] = nums[pos], nums[i]
            pos++
        }
    }
}