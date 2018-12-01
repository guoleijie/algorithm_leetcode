package main

import "fmt"
import "sort"

func minincrement(A []int) {
	sort.Ints(A)
	resp := 0
	for i := 1; i < len(A); i = i + 1 {
		if A[i-1] >= A[i] {
			resp = resp + A[i-1] - A[i] + 1
			A[i] = A[i-1] + 1
		}
	}
	fmt.Println(A)
	fmt.Println(resp)
}
func main() {
	var A = []int{2, 6, 10, 6, 8, 6, 9, 4, 3, 3, 2, 2, 3, 3, 5, 6}
	minincrement(A)
}
