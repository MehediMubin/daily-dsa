// Problem: Maximum Number of Balloons 
// Link: https://leetcode.com/problems/maximum-number-of-balloons/description/
// Time Complexity: O(n)
// Space Complexity: O(1)

func min(numbers ...int) int {
  ans := numbers[0]
  for _, number := range numbers {
    if number < ans {
      ans = number
    }
  }
  return ans
}

func maxNumberOfBalloons(text string) int {
    mp := map[string]int{}
    for _, v := range text {
      mp[string(v)]++
    }

    ans := min(mp["b"], mp["a"], mp["l"] / 2, mp["o"] / 2, mp["n"])
    return ans
}