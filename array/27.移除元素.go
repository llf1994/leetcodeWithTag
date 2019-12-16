package main
//最终答案
func removeElement__(nums []int, val int) int {
	left, right := 0, len(nums) - 1

	//当left与right相遇时停止
	for left <= right {
		//left从左往右，捕获nums中与val相等的元素
		for left <= right && nums[left] != val {
			left++
		}
		//right从右往左，捕获nums中的普通元素
		for left <= right && nums[right] == val{
			right--
		}
		//未越界则交换
		if left < right{
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return left
}

//题解
func removeElement_(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == val{
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return right + 1
}


//第一次通过答案
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums) - 1

	//若nums中不存在val，直接返回
	have := false
	for _, v := range nums{
		if v == val{
			have = true
			break
		}
	}
	if have == false{
		return len(nums)
	}

	//当left与right相遇时停止
	for left < right {
		//left从左往右，捕获nums中与val相等的元素
		for left < right && nums[left] != val {
			left++
		}
		//right从右往左，捕获nums中的普通元素
		for left < right && nums[right] == val{
			right--
		}
		//未越界则交换
		if left < right{
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return left
}