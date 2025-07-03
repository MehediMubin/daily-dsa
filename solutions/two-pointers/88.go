// Problem: Merge Sorted Array
// Link: https://leetcode.com/problems/merge-sorted-array/description/
// Time Complexity: O(m + n)
// Space Complexity: O(m + n) 

func merge(nums1 []int, m int, nums2 []int, n int)  {
    arr := make([]int, m + n)
    l, r, i := 0, 0, 0
    for l < m && r < n {
      if nums1[l] < nums2[r] {
        arr[i] = nums1[l]
        l++
      } else {
        arr[i] = nums2[r]
        r++
      }
      i++
    }

    for l < m {
      arr[i] = nums1[l]
      i, l = i + 1, l + 1
    }

    for r < n {
      arr[i] = nums2[r]
      i, r = i + 1, r + 1
    }

    copy(nums1, arr)
}