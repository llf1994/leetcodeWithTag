package main

//简化版
func removeDuplicates2_(nums []int) int {
	length := len(nums)
	if length <= 2{return length}

	slow := 2
	for fast:=2; fast < length; fast++{
		if nums[fast] != nums[slow-2]{
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

//第一次通过答案
func removeDuplicates2(nums []int) int {
	length := len(nums)
	//当数组长度小于等于2时满足题意，直接返回
	if length <= 2{
		return length
	}

	//前两个元素符合题意，从第三个元素开始
	//慢指针slow指向待修改的元素，快指针fast遍历数组
	slow, fast := 2, 2

	for fast < length{
		//当s[fast]==s[slow-2]时，因原数组有序，所以此时s[fast]==s[slow-1]==s[slow-2]，不合题意，往下遍历寻找新元素
		if nums[fast] == nums[slow-2]{
			fast++
			//当s[fast]!=s[slow-2]时，找到符合要求的新元素，将新元素写入s[slow]，并把slow和fast指针均后移一位
		} else{
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	return slow
}