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