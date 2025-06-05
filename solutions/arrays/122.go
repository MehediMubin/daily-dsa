// Problem: Best Time to Buy and Sell Stock 2
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
// Time Complexity: O(n)
// Space Complexity: O(1)

func maxProfit(prices []int) int {
    current := prices[0]
    ans := 0
    for _, price := range(prices) {
        if price > current {
            ans += price - current
        }
        current = price
    }

    return ans
}