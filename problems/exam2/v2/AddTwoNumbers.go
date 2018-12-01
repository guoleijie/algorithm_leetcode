package main

import "fmt"

//SUCCESS
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	var resultStart *ListNode
	resultStart = result

	carry := 0
	for ; ; {
		var lv1 int
		var lv2 int
		if l1 == nil {
			lv1 = 0
		} else {
			lv1 = l1.Val
		}
		if l2 == nil {
			lv2 = 0
		} else {
			lv2 = l2.Val
		}
		result.Val = (lv1 + lv2 + carry) % 10
		carry = (lv1 + lv2 + carry) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		result.Next = &ListNode{}
		result = result.Next
	}
	return resultStart
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
	a := []int{5}
	b := []int{5}
	//b := []int{5, 6, 4}
	var n ListNode
	for i := len(a) - 1; i >= 0; i-- {
		n = ListNode{a[i], &ListNode{n.Val, n.Next}}
	}
	var m ListNode
	for i := len(b) - 1; i >= 0; i-- {
		m = ListNode{b[i], &ListNode{m.Val, m.Next}}
	}
	result := addTwoNumbers(&n, &m)
	for ; result != nil; {
		fmt.Println(result.Val)
		result = result.Next
	}
}
