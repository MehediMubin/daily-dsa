// Problem: Longest Common Prefix
// Link: https://leetcode.com/problems/longest-common-prefix/description/
// Time Complexity: O(n * m^2), where n is the number of strings and m is the maximum possible length of the common prefix
// Space Complexity: O(1)

import "strings"

func longestCommonPrefix(strs []string) string {
    s := strs[0]
    n := len(strs)

    for len(s) > 0 {
        for i, str := range strs {
            if !strings.HasPrefix(str, s) {
                break
            }
            if i == n - 1 {
                return s
            }
        }
        s = s[:len(s) - 1]
    }
    return s
}