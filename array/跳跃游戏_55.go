package main

//题解中提到动态规划（自顶向下和自底向上两种），有空实现一下
//ps:原来这就是贪心算法，这里使用了从后往前的贪心算法，也可以实现从前往后（不断更新能跳到最远的距离）

//优化for循环，同样为o(n)
func canJump(nums []int) bool {
	//快慢指针，慢指针指向能到达终点的位置(包括终点自身)，快指针寻找能到达慢指针的位置
	slow := len(nums) - 1
	//当fast==0时结束
	for fast := slow - 1; fast >= 0; fast-- {
		//当fast + nums[fast] >= slow时，fast指向的位置能到达slow，此时重设快慢指针
		if fast+nums[fast] >= slow {
			slow = fast
		}
		//如果fast指向的位置不能到达slow，继续往前寻找
	}
	//如果能够到达，此时slow指针指向第一个位置
	return slow == 0
}

//第一次答案，时间复杂度为o(n)
func canJump_first(nums []int) bool {
	//快慢指针，慢指针指向能到达终点的位置(包括终点自身)，快指针寻找能到达慢指针的位置
	slow, fast := len(nums)-1, len(nums)-2
	//当fast==0时结束
	for fast >= 0 {
		//当fast + nums[fast] >= slow时，fast指向的位置能到达slow，此时重设快慢指针
		if fast+nums[fast] >= slow {
			slow, fast = fast, fast-1
		} else {
			//如果fast指向的位置不能到达slow，继续往前寻找
			fast--
		}
	}
	//如果能够到达，此时slow指针指向第一个位置
	return slow == 0
}
