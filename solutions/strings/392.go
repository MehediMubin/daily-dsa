// Problem: Is Subsequence
// Link: https://leetcode.com/problems/is-subsequence/description/
// Time Complexity: O(n), where n is the length of the string t
// Space Complexity: O(1)


func isSubsequence(s string, t string) bool {
    sLen := len(s)
    idx := 0
    if sLen == 0 {
        return true
    }
    for i := 0; i < len(t); i++ {
        if s[idx] == t[i] {
            idx++
        }
        if idx == sLen {
            return true
        }
    }
    return false
}