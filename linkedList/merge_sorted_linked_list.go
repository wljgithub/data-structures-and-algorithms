package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果 l1 为空，返回l2
	if l1 == nil {
		return l2
	}
	// 如果 l2 为空，返回l1
	if l2 == nil {
		return l1
	}
	// 比较两个链表的元素，取出大的元素，然后将剩下的递归返回
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
func makeLinkedList(eles []int) *ListNode {
	if len(eles) <= 0 {
		return nil
	}
	head := &ListNode{
		Val: eles[0],
	}

	current := head
	for _, v := range eles[1:] {
		addToEnd(current, v)
	}
	return head
}
func addToEnd(head *ListNode, val int) *ListNode {
	ele := &ListNode{Val: val}
	if head == nil {
		return ele
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = ele
	return head

}
func main() {

	node1 := makeLinkedList([]int{1, 2, 4})
	node2 := makeLinkedList([]int{1, 3, 4})

	mergeTwoLists(node1, node2)
}
