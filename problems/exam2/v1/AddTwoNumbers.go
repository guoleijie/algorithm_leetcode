package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := 0
	var resultStart ListNode
	var resultEnd *ListNode
	if (l1.Val+l2.Val)/10 == 1 {
		resultStart = ListNode{Next: &ListNode{}}
		resultEnd = resultStart.Next
	} else {
		resultStart = ListNode{}
		resultEnd = &resultStart
	}
	resultStart.Val = temp + (l1.Val+l2.Val)/10
	temp = (l1.Val + l2.Val) % 10
	l1 = l1.Next
	l2 = l2.Next
	for ; l1 != nil || l2 != nil; {

	}
	return &resultStart
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */
func main() {
	a := []int{2, 4, 3}
	b := []int{5, 6, 4}
	//b := []int{5, 6, 4}
	var n ListNode
	for i := len(a) - 1; i >= 0; i-- {
		n = ListNode{a[i], &ListNode{n.Val, n.Next}}
	}
	var m ListNode
	for i := len(b) - 1; i >= 0; i-- {
		m = ListNode{b[i], &ListNode{m.Val, m.Next}}
	}
	//fmt.Print(temp)

	/*for ; n.Next != nil; {
		fmt.Println(n.Val)
		n = *n.Next
	}

	for ; m.Next != nil; {
		fmt.Println(m.Val)
		m = *m.Next
	}*/
	result := addTwoNumbers(&n, &m)
	for ; result.Next != nil; {
		fmt.Println(result.Val)
		result = result.Next
	}
}
