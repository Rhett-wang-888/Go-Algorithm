/**
 * @description: TODO 两个有序链表排序
 * @author rhett.wang
 * @date 2021/11/1 11:00
 * @version 1.0
 */
package main

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
