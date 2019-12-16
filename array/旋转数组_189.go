package main

//三次反转代码最简洁，但最难想到

//题解，环状替换，每次移动k个位置，用计数器count保证所有元素均移动
func rotate(nums []int, k int) {
	//计数器count每次移动元素时自增，指针i为初始移动元素下标
	count, i, length, k := 0, 0, len(nums), k % len(nums)
	//count==length时结束，此时所有元素都完成移动
	for count < length {
		tmp := nums[i]
		//当j==i时移动结束
		for j := i + k; j != i; j = (j + k) % length {
			nums[j], tmp = tmp, nums[j]
			count++
		}
		nums[i] = tmp

		//tmp也移动了1次
		count++

		//若count!=length，有元素未移动，i自增继续移动
		i++
	}
}

//题解，三次反转，等价于数组向右移动k
func rotate1(nums []int, k int) {
	length, k := len(nums), k%len(nums)
	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	//反转数组
	reverse(nums, 0, length-1)
	//反转前k个元素，k%n 个尾部元素被移动到头部
	reverse(nums, 0, k-1)
	//反转后 n-k 个数字
	reverse(nums, k, length-1)
}

//暴力法
func rotate_(nums []int, k int) {
	length := len(nums)
	//如果k>=len(nums)，化简
	if k >= length {
		k %= length
	}

	//移动k次
	for k > 0 {
		//向右移动1个位置，需一个指针和一个tmp值
		//将第一个值放入tmp中
		tmp := nums[0]
		//每次遍历交换tmp与指针值
		for i := 1; i < length; i++ {
			nums[i], tmp = tmp, nums[i]
		}
		//将tmp放入第一个值中
		nums[0] = tmp
		k--
	}
}
