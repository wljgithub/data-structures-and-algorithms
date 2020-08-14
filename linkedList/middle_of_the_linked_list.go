package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	walk := head
	run := head

	for run!=nil && run.Next != nil {

		walk = walk.Next
		run = run.Next.Next
	}
	return walk
}