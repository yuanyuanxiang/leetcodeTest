package leetcodeTest

import (
	"fmt"
	"testing"
)

func Test0082(t *testing.T) {
	fmt.Println(deleteDuplicates(NewListNode([]int{1, 2, 3, 3, 4, 4, 5})))
	fmt.Println(deleteDuplicates(NewListNode([]int{1, 2, 3, 3, 4, 4, 5, 5, 5})))
	fmt.Println(deleteDuplicates(NewListNode([]int{1, 1, 1, 2, 3})))
	fmt.Println(deleteDuplicates(NewListNode([]int{1, 1, 2, 2})))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
【82.删除排序链表中的重复元素 II】
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
返回同样按升序排列的结果链表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
*/
func deleteDuplicates(head *ListNode) *ListNode {
	var prev *ListNode
	var cur = head
	if cur == nil {
		return nil
	}
	var next = cur.Next
	for next != nil {
		if cur.Val != next.Val {
			prev = cur
			cur = next
			next = next.Next
			continue
		}
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		if prev == nil {
			head = next
		} else {
			prev.Next = next
		}
		if next == nil {
			break
		}
		cur = next
		next = cur.Next
	}
	return head
}
