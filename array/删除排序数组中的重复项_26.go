package main

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j < len(nums){
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i+1
}
