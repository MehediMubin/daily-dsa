// Problem: Reorganize String
// Link: https://leetcode.com/problems/reorganize-string/description/
// Time Complexity: O(n + k log k)
// Space Complexity: O(n)

func reorganizeString(s string) string {
  mp := map[byte]int{}
  for i := 0; i < len(s); i++ {
    mp[s[i]]++
  }

  sortedChar := []byte{}
  for key := range mp {
    sortedChar = append(sortedChar, key)
  }
  sort.Slice(sortedChar, func(i, j int) bool {
    return mp[sortedChar[i]] > mp[sortedChar[j]]
  })

  res := make([]string, len(s))
  index := 0
  for i := 0; i < len(sortedChar); i++ {
    charFreq := mp[sortedChar[i]]
    if charFreq > (len(s) + 1) / 2 {
      return ""
    }

    for range charFreq {
      if index >= len(s) {
        index = 1
      }
      res[index] = string(sortedChar[i])
      index += 2
    }
  }

  return strings.Join(res, "")
}