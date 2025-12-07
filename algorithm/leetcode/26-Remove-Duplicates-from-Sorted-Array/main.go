package main

func removeDuplicates(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1
	fast := 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
