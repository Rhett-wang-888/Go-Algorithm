/**
 * @description: TODO 链表归并排序  自顶向上
 * @author rhett.wang
 * @date 2021/10/25 14:20
 * @version 1.0
 */
package main

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, curr := dummyHead, dummyHead.Next

		for curr != nil {
			head1 := curr
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			curr.Next = nil
			curr = head2
			for i := 1; i < subLength && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}
			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil
			}
			prev.Next = merge2(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = next
		}
	}

	return dummyHead.Next
}
func merge2(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}

	return dummyHead.Next
}
