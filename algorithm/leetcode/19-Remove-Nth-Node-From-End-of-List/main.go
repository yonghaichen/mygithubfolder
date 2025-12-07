package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	first := head
	for i := 0; i < n; i++ {
		first = first.Next
	}

	dummy := &ListNode{0, head}
	second := dummy
	for ; first != nil; first = first.Next {
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}
