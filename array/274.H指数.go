package main

import (
	"sort"
)


//使用计数排序，时间复杂度为o(n)，空间复杂度为o(n)
func hIndex__(citations []int) (h int) {
	//由于h<=len，遍历数组，将引用数大于论文数的值改成论文数
	length := len(citations)
	for i, v := range citations {
		if v > length {
			citations[i] = length
		}
	}

	//使用计数排序，准备length+1个桶
	buckets := make([]int, length+1)
	//将数字全部存入桶中
	for _, v := range citations {
		buckets[v]++
	}
	//逆序从桶中取出数字，得到降序数组
	for i, j := 0, 0; i < length+1; i++{
		for buckets[length-i] > 0 {
			citations[j] = length - i
			buckets[length-i]--
			j++
		}
	}

	//遍历降序数组，当s[i]>=i+1时，符合题意
	for i, v := range citations {
		//如果v<i+1，退出循环，此时h的值为上一轮循环的 索引i 对应的 论文数i+1
		if v < i+1 {
			break
		}
		h = i + 1
	}
	return h
}

//语言内置排序，时间复杂度为o(nlogn)
func hIndex_(citations []int) (h int) {
	//降序排序，使用sort.Reverse更改默认升序的排序行为（改变Less方法）
	//sort.Sort方法整体框架是快排（同时使用了堆排序和希尔排序），时间复杂度为o(n logn)
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	//此时索引i对应的论文数量为i+1
	//遍历降序数组，当s[i]>=i+1时，符合题意
	for i, v := range citations {
		//如果v<i+1，退出循环，此时h的值为上一轮循环的 索引i 对应的 论文数i+1
		if v < i+1 {
			break
		}
		h = i + 1
	}
	return h
}
