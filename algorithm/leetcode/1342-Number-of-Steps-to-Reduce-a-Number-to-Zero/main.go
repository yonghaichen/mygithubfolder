package main

func numberOfStepsBa(num int) int {

	var count int

	for num > 0 {
		count += (num & 1) + 1
		num >>= 1
	}

	if count-1 > 0 {
		return count - 1
	}
	return 0
}
