package main

import "fmt"

func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if val, ok := data[temp]; ok {
			return []int{val, i}
		}
		data[nums[i]] = i
	}
	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	resp := twoSum(arr, target)
	fmt.Println(resp)
}
