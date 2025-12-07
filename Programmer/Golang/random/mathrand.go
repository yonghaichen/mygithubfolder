package random

import (
	"bytes"
	"math/rand"
	"time"
)

func GenerateRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start))
		exist := false
		for _, v := range nums {
			if (start + num) == v {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, start+num)
		}
	}
	return nums
}

func GenerateRandomString(len int64) string {
	strBytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := []byte{}
	var i int64
	for i = 0; i < len; i++ {
		idx := r.Intn(bytes.Count(strBytes, nil) - 1)
		temp := strBytes[idx]
		result = append(result, temp)
	}
	return string(result)
}
