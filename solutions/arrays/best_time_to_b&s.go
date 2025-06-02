// Problem: Best Time to Buy and Sell Stock
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Time Complexity: O(n)
// Space Complexity: O(1) 

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxProfit(prices []int) int {
    mxProfit := 0
    current := prices[0]
    for _, price := range(prices) {
        if price < current {
            current = price
        } else {
            mxProfit = max(mxProfit, price - current)
        }
    }
    return mxProfit
}