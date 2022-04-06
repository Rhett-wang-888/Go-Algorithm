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

func mergeTwoListS4(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{0, nil}

	prev := prehead

	for {
		if l1 == nil {
			break
		}
		if l2 == nil {
			break
		}

		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next

	}
	if l1 == nil {
		prev.Next = l2
	}
	if l2 == nil {
		prev.Next = l1
	}
	return prehead.Next
}

func mergeTwoList5(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	cur := new(ListNode)
	res := cur

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return res.Next
}

func mergeTwoList6(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l3Fast *ListNode = nil
	var l3Last *ListNode = nil
	var newNode *ListNode = nil

	for true {
		if l1.Val > l2.Val {
			newNode = l2
			l2 = l2.Next
		} else {
			newNode = l1
			l1 = l1.Next
		}

		if l3Fast == nil {
			l3Fast = newNode
			l3Last = newNode
		} else {
			l3Last.Next = newNode
			l3Last = newNode
		}

		if l1 == nil {
			l3Fast = newNode
			l3Last = newNode
		} else {
			l3Last.Next = newNode
			l3Last = newNode
		}
		if l1 == nil {
			l3Last.Next = l2
			break
		} else if l2 == nil {
			l3Last.Next = l1
			break
		}
	}

	return l3Fast
}

func mergeTwoList7(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoList7(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList7(l1, l2.Next)
		return l2
	}

}
func mergeTwoList8(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}

	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return dummyHead.Next
}
func mergeTwoList11(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	cur := new(ListNode)
	res := cur

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return res.Next
}
func mergeTwoList13(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cursor := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}
	if list1 == nil {
		cursor.Next = list2
	} else {
		cursor.Next = list1
	}

	return dummyHead.Next
}

func mergeTwoList12(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoList12(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoList12(list1, list2.Next)
		return list2
	}
}
