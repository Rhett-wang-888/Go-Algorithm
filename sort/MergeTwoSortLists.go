/**
 * @description: TODO 两个有序链表排序
 * @author rhett.wang
 * @date 2021/11/1 11:00
 * @version 1.0
 */
package main

func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	var protect = &ListNode{0, l1}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	end := getEnd(l1)
	end.Next = l2

	return sortList3(protect.Next)
}

func sortList3(head *ListNode) *ListNode {
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

			prev.Next = merge3(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = next
		}
	}
	return dummyHead.Next
}

func getEnd(l1 *ListNode) *ListNode {
	k := 100
	head := l1
	for head.Next != nil {
		k--
		head = head.Next
	}

	return head
}
func merge3(head1, head2 *ListNode) *ListNode {
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

/////////////////////////////////////////////////

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				curr.Next = l1

				l1 = l1.Next
			} else {
				curr.Next = l2
				l2 = l2.Next
			}
			curr = curr.Next
		} else if l1 != nil {
			curr.Next = l1
			break
		} else {
			curr.Next = l2
			break
		}
	}
	return head.Next
}

//递归
func mergeTwoList1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoList1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList1(l1, l2.Next)
		return l2
	}

	return nil
}

//迭代
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}
