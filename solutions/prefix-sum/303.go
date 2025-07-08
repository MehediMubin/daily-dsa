// Problem: Range Sum Query - Immutable
// Link: https://leetcode.com/problems/range-sum-query-immutable/
// Time Complexity: O(n)
// Space Complexity: O(1)

type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    var arr NumArray
    arr.nums = nums
    return arr
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for i := left; i <= right; i++ {
      sum += this.nums[i]
    }
    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */