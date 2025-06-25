// Problem: Ransom Note
// Link: https://leetcode.com/problems/ransom-note/description/
// Time Complexity: O(m + n)
// Space Complexity: O(1)

func canConstruct(ransomNote string, magazine string) bool {
    mp := map[byte]int{}
    for _, value := range magazine {
      mp[byte(value)]++
    }

    for _, value := range ransomNote {
      if mp[byte(value)] == 0 {
        return false
      }
      mp[byte(value)]--
    }
    return true
}