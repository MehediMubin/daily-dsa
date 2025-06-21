// Problem: Single Number III
// Link: https://leetcode.com/problems/single-number-iii/
// Time Complexity: O(n)
// Space Complexity: O(n)

func singleNumber(nums []int) []int {
    mp := map[int]int{}
    for _, num := range nums {
      mp[num]++
    }
    
    arr := []int{}
    for key, value := range mp {
      if value == 1 {
        arr = append(arr, key)
      }
    }
    return arr
}