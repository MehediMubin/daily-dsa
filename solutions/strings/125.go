// Problem: Valid Palindrome 
// Link: https://leetcode.com/problems/valid-palindrome/
// Time Complexity: O(n)
// Space Complexity: O(n)

import "unicode"

func isPalindrome(s string) bool {
    var newString string
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            newString += string(unicode.ToLower(char))
        }
    }

    s = newString
    l, r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++; r--
    }
    return true
}