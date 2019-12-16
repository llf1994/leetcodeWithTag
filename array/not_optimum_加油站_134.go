package main

//题解
func canCompleteCircuit(gas []int, cost []int) int {
	//o(n)复杂度的算法，待完成
	return -1
}

//暴力法，时间复杂度为o(n^2)
func canCompleteCircuit_(gas []int, cost []int) int {
	length := len(gas)
	for i:=0; i<length;i++{
		//到达下一个地点时，油箱里剩下remain升油
		//j为当前站点，此时已经开过一个加油站(gas[i] - cost[i])，所以j=i+1
		j, remain := i+1, gas[i] - cost[i]
		for remain >= 0 {
			// j到达边界时重置
			if j == length {j = 0}

			//检查是否返回
			if i == j {return i}

			//开往下一个站点
			remain = remain + gas[j] - cost[j]
			j++
		}
	}
	return -1
}
