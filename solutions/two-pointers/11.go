// Problem: Container with Most Water
// Link: https://leetcode.com/problems/container-with-most-water/description/
// Time Complexity: O(n)
// Space Complexity: O(1)

func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	ans := 0
	for l < r {
		ans = max(ans, (r - l) * min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}