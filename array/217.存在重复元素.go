package main

//哈希表，时间和空间复杂度均为o(n)
func containsDuplicate(nums []int) bool {
	hashmap := map[int]bool{}
	for _, v := range nums{
		if _, ok := hashmap[v]; ok{
			return true
		}
		hashmap[v] = true
	}
	return false
}