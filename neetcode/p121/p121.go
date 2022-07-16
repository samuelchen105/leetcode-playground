package problem

// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	ret := 0
	l := 0
	for r := 1; r < len(prices); r++ {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > ret {
				ret = profit
			}
		} else {
			l = r
		}
	}
	return ret
}
