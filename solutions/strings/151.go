// Problem: Reverse words in a String
// Link: https://leetcode.com/problems/reverse-words-in-a-string/description/
// Time Complexity: O(n)
// Space Complexity: O(n)

func populate(s string) []string {
  var str strings.Builder
  arr := []string{}

  for i := range s {
    if string(s[i]) == " " {
      if len(str.String()) > 0 {
        arr = append(arr, str.String())
        str.Reset()
      }
    } else {
        str.WriteString(string(s[i]))
      }
  }

  if len(str.String()) > 0 {
    arr = append(arr, str.String())
  }

  return arr
}

func reverseWords(s string) string {
  arr := populate(s)
  var builder strings.Builder

  for i := len(arr) - 1; i >= 0; i-- {
    builder.WriteString(arr[i])
    if i > 0 {
      builder.WriteString(" ")
    }
  }
  return builder.String()
}