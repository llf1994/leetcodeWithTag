package main

//题解
//基于桶排序的变式，利用了索引与值的映射关系（索引i上的值应为i+1）
//容易推断首次缺失的正数一定小于或等于 len+1
func firstMissingPositive(nums []int) int {
	length := len(nums)
	//两次遍历，首次遍历按索引逐个检查，保证索引上的值满足以下情况之一
	//(1)正确（符合基于索引的映射关系）
	//(2)越界（该索引不在数组的区间范围内）
	//(3)该索引值对应的索引上已有正确的值（数组中有重复的值）
	for i := range nums {
		//索引值nums[i]>0且其对应的索引nums[i]-1未越界，即该值为数组nums中的索引（nums[i] > 0 && nums[i] <= length）
		//如果未越界，且该索引值对应的索引上的值不正确（nums[i] != nums[nums[i]-1]）
		for nums[i] > 0 && nums[i] <= length && nums[i] != nums[nums[i]-1] {
			//当索引i对应的值v不等于索引i+1时，将该值放入对应的索引中（此时对应索引上的值不正确）
			if v := nums[i]; v != i+1 {
				nums[i], nums[v-1] = nums[v-1], nums[i]
			}
			//在下一次循环中检查交换后该索引值是否需要继续放置
		}
	}

	//第二次寻找不符合映射关系的元素
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	//若寻找失败，则首次缺失的正数为n+1
	return length + 1
}
