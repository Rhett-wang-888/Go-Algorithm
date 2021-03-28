/**
 * @description: TODO 二叉搜索树
 * @author rhett.wang
 * @date 2021/3/28 13:33
 * @version 1.0
 */
package main

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}
type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inorder(root)
	return
}
func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.right)
}
func (it *BSTIterator) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}
func (it *BSTIterator) HasNext() bool {
	return len(it.arr) > 0
}
