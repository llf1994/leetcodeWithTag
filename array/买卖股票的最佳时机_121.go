package main

//使用动态规划的系列题目，这题没有用到

//题解，多记录一个最小值,o(n)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {return 0}

	maxPro, minPrice := 0, prices[0]
	for _, v := range prices{
		//如果当前卖出得到的利润大于之前利润，更新利润
		if v - minPrice > maxPro{
			maxPro = v - minPrice
		}
		//如果当前值比最小值小，更新最小值
		if v < minPrice{
			minPrice = v
		}
	}
	return maxPro
}

//第一次通过答案，暴力法，o(n^2)
func maxProfit_(prices []int) int {
	maxPro := 0
	for i, v := range prices{
		for j:=i+1;j<len(prices);j++{
			//如果价格大于买入价格，且利润比之前的高，刷新利润
			if prices[j] - v > maxPro {
				maxPro = prices[j] - v
			}
		}
	}
	return maxPro
}