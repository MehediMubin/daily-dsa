// Problem: Reverse Bits 
// Link: https://leetcode.com/problems/reverse-bits/description/
// Time Complexity: O(1)
// Space Complexity: O(1)

func reverseBits(num uint32) uint32 {
    var result uint32
    for i := 0; i < 32; i++ {
      result <<= 1
      result |= num & 1
      num >>= 1
    }
    return result
}