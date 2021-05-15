/**
 * @description: TODO 指定位置翻转链表
 * @author rhett.wang
 * @date 2021/5/15 16:22
 * @version 1.0
 */
package main

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}
func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	curr := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	reverseLinkedList(leftNode)
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}
func reverseBetween1(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
