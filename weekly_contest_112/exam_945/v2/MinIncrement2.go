package main

import (
	"fmt"
	"sort"
)

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	count := 0
	//循环所有元素
	for index, num := range A {
		tempcount, _ := getIncrementCount(A, num, index, 0, -1, false)
		count = count + tempcount
	}
	return count
}

func getIncrementCount(A []int, a int, index int, count int, x int, needmoved bool) (int, int) {
	//出口1：
	//* 如果索引到最后位置，不需要再移动，但是需要与x进行判断,如果a<=x。a移动一个位置
	if len(A) == index+1 {
		if needmoved || a <= x {
			return count + IntMax(x, a) - a, x
		}
		return count, x
	}
	b := A[index+1]
	//出口2: 如果完全不需要移动
	if !needmoved && b > a && x < a {
		return count, x
	}
	//出口3：
	//* 如果差值大于1，说明中间存在中空位置，移动a一个位置就可以了。并记录下a最后的值并赋值给x
	if needmoved && b-IntMax(x, a) > 1 {
		return count + (IntMax(a, x) - a) + 1, IntMax(a, x) + 1
	}
	// * 如果差值等于1，说明值相邻，需要进行迭代判断b的下个索引与b的位置差异。
	return getIncrementCount(A, b, index+1, count+1, x, true)
}
func IntMax(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	var A = []int{2,2,3,4,5,10}
	count := minIncrementForUnique(A)
	fmt.Print(",num:", count)
}
