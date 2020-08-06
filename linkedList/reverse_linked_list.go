package main


//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil
	return root
}

func main(){
	node4:=&ListNode{
		Val:  4,
		Next: nil,
	}
	node3:=&ListNode{
		Val:  3,
		Next: node4,
	}

	node2:=&ListNode{
		Val:  2,
		Next: node3,
	}
	node1:=&ListNode{
		Val:  1,
		Next: node2,
	}
	reverseList(node1)

}
