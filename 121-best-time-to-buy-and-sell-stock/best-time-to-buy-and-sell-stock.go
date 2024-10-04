func maxProfit(prices []int) int {
    min := prices[0]
    max := 0

    for _, p := range prices {
        if p < min {
            min = p
        }
        profit := p - min
        if profit > max {
            max = profit
        }
    }

    return max
}