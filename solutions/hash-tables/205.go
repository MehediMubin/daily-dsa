// Problem: Isomorphic Strings 
// Link: https://leetcode.com/problems/isomorphic-strings/
// Time Complexity: O(n)
// Space Complexity: O(n)

func isIsomorphic(s string, t string) bool {
    mp1 := map[byte]byte{}
    mp2 := map[byte]byte{}
    for i := 0; i < len(s); i++ {
      if mp1[s[i]] == 0 && mp2[t[i]] == 0{
        mp1[s[i]] = t[i]
        mp2[t[i]] = s[i]
      }
    }

    for i := 0; i < len(s); i++ {
      if mp1[s[i]] != t[i] {
        return false
      }
    }
    return true
}