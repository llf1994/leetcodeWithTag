package main

import "math"

//题解
func increasingTriplet(nums []int) bool {
	//将i、j初始值设为max，其中i为最小值，j为次小值
	i, j := math.MaxInt64, math.MaxInt64

	//遍历数组
	for _, v := range nums{
		//首先将v与i比较，如果v<=i，说明i不是最小值，则i=v
		if v <= i {
			i = v
		//之后将v与j比较，此时v>i && j>i；如果v<=j，说明在i、j、v中j不是次小值，则j=v
		} else if v <= j {
			j = v
		//如果进入该分支，说明v>i && v>j，存在符合条件的子序列
		} else {
			return true
			//此时i、j、k的子序列不一定是递增序列，因为递增子序列中i的值可能为i更新前的值；
			//但依然符合题意，以下简单证明：
			//(1)假设i0为i更新前的值，此时j>i0;
			//(2)而只有i<i0时，i才会被更新，即i0>i;
			//(3)由(1)(2)得j>i，j>i0，无论如何j一定为次小值
			//而进入该分支的条件为v>j，则v>j>i和v>j>i0均成立，更新最小值i来改变递增子序列时，不影响最终判断
		}
	}
	return false
}
