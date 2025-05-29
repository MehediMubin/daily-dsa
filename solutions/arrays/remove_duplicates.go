func removeDuplicates(nums []int) int {
    k := 1
    pos := 0
    for i := 1; i < len(nums); i++ {
        if nums[pos] != nums[i] {
            pos++
            nums[pos] = nums[i]
            k++
        }
    }
    return k
}