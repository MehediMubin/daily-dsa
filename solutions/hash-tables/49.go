// Problem: Group Anagrams
// Link: https://leetcode.com/problems/group-anagrams/description/
// Time Complexity: O(n * k log k)
// Space Complexity: O(n * k)

func groupAnagrams(strs []string) [][]string {
  mp := map[string][]string{}
  for i := 0; i < len(strs); i++ {
    str := []byte(strs[i])
    sort.Slice(str, func (i, j int) bool {
      return str[i] < str[j]
    })

    mp[string(str)] = append(mp[string(str)], strs[i])
  }

  ans := [][]string{}
  for _, val := range mp {
    ans = append(ans, val)
  }

  return ans
}