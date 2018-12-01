package v1

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == temp {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	resp := twoSum(arr, target)
	fmt.Println(resp)
}
