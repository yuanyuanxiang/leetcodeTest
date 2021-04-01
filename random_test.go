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

func Test0083(t *testing.T) {
	fmt.Println(deleteDuplicates83(NewListNode([]int{1, 1})))
	fmt.Println(deleteDuplicates83(NewListNode([]int{1, 1, 1})))
	fmt.Println(deleteDuplicates83(NewListNode([]int{1, 1, 2})))
	fmt.Println(deleteDuplicates83(NewListNode([]int{1, 1, 2, 3, 3})))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
【83. 删除排序链表中的重复元素】
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
返回同样按升序排列的结果链表。

示例 1：

输入：head = [1,1,2]
输出：[1,2]

示例 2：

输入：head = [1,1,2,3,3]
输出：[1,2,3]


提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteDuplicates83(head *ListNode) *ListNode {
	var cur = head
	if cur == nil {
		return nil
	}
	var next = cur.Next
	for next != nil {
		if cur.Val != next.Val {
			cur = next
			next = cur.Next
			continue
		}
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		cur.Next = next
		if next == nil {
			break
		}
		cur = next
		next = cur.Next
	}
	return head
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0061(t *testing.T) {
	fmt.Println(rotateRight(NewListNode([]int{1, 2, 3, 4, 5}), 2))
	fmt.Println(rotateRight(NewListNode([]int{0, 1, 2}), 4))
	fmt.Println(rotateRight(NewListNode([]int{1, 2}), 0))
	fmt.Println(rotateRight(NewListNode([]int{1, 2}), 1))
	fmt.Println(rotateRight(NewListNode([]int{1, 2}), 2))
	fmt.Println(rotateRight(NewListNode([]int{1, 2, 3}), 2000000000))
}

/*
【61.旋转链表】
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func rotateRight(head *ListNode, k int) *ListNode {
	cur := head
	if cur == nil || k <= 0 {
		return cur
	}
	next := cur.Next
	length := 1
	for next != nil {
		cur = next
		next = cur.Next
		length++
	}
	offset := length - k%length
	if offset == 0 {
		return head
	}
	cur.Next = head
	cur = cur.Next
	for i, n := 0, offset-1; i < n; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil

	return head
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0173(t *testing.T) {
	var test = Constructor(&TreeNode{Val: 10, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 10}})
	fmt.Println(test)
	var n3 = TreeNode{Val: 3}
	var n9 = TreeNode{Val: 9}
	var n20 = TreeNode{Val: 20}
	var n15 = TreeNode{Val: 15, Left: &n9, Right: &n20}
	var n7 = TreeNode{Val: 7, Left: &n3, Right: &n15}
	var bSTIterator = Constructor(&n7)
	fmt.Println(bSTIterator)
	fmt.Println(bSTIterator.Next())    // 返回 3
	fmt.Println(bSTIterator.Next())    // 返回 7
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 9
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 15
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 20
	fmt.Println(bSTIterator.HasNext()) // 返回 False
}

/*
【173. 二叉搜索树迭代器】
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。
指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。

你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

示例：

输入
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
输出
[null, 3, 7, true, 9, true, 15, true, 20, false]

解释
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False

提示：
树中节点的数目在范围 [1, 10^5] 内
0 <= Node.val <= 10^6
最多调用 10^5 次 hasNext 和 next 操作

进阶：
你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。其中 h 是树的高度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator 内存消耗较大.
type BSTIterator struct {
	Index   int
	Length  int
	Results []int
}

// Constructor 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。
// 指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
func Constructor(root *TreeNode) BSTIterator {
	it := recursive(root)
	return BSTIterator{Index: 0, Results: it, Length: len(it)}
}

func recursive(root *TreeNode) []int {
	var left, right []int
	if root.Left != nil {
		left = recursive(root.Left)
	}
	if root.Right != nil {
		right = recursive(root.Right)
	}
	left = append(left, root.Val)
	return append(left, right...)
}

// Next 将指针向右移动，然后返回指针处的数字。
func (this *BSTIterator) Next() int {
	if this.Index < this.Length {
		r := this.Results[this.Index]
		this.Index++
		return r
	}
	return -1
}

// HasNext 如果向指针右侧遍历存在数字，则返回 true；否则返回 false。
func (this *BSTIterator) HasNext() bool {
	return this.Index < this.Length
}

func Test0173_2(t *testing.T) {
	var n3 = TreeNode{Val: 3}
	var n9 = TreeNode{Val: 9}
	var n20 = TreeNode{Val: 20}
	var n15 = TreeNode{Val: 15, Left: &n9, Right: &n20}
	var n7 = TreeNode{Val: 7, Left: &n3, Right: &n15}
	var bSTIterator = Constructor2(&n7)
	fmt.Println(bSTIterator)
	fmt.Println(bSTIterator.Next())    // 返回 3
	fmt.Println(bSTIterator.Next())    // 返回 7
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 9
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 15
	fmt.Println(bSTIterator.HasNext()) // 返回 True
	fmt.Println(bSTIterator.Next())    // 返回 20
	fmt.Println(bSTIterator.HasNext()) // 返回 False
}

/* BSTIterator2 节省内存的结构. */
type BSTIterator2 struct {
	Length  int
	Results []*TreeNode
}

// Constructor 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。
// 指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
func Constructor2(root *TreeNode) BSTIterator2 {
	var r = BSTIterator2{Results: make([]*TreeNode, 0, 16)}
	for root != nil {
		r.Results = append(r.Results, root)
		root = root.Left
	}
	r.Length = len(r.Results)
	return r
}

// Next 将指针向右移动，然后返回指针处的数字。
func (this *BSTIterator2) Next() int {
	if len(this.Results) == 0 {
		return -1
	}
	var cur = this.Results[this.Length-1]
	this.Results = this.Results[:this.Length-1]
	var r = cur.Val
	if cur.Right != nil {
		this.Results = append(this.Results, cur.Right)
		for root := cur.Right.Left; root != nil; root = root.Left {
			this.Results = append(this.Results, root)
		}
	}
	this.Length = len(this.Results)

	return r
}

// HasNext 如果向指针右侧遍历存在数字，则返回 true；否则返回 false。
func (this *BSTIterator2) HasNext() bool {
	return this.Length != 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0190(t *testing.T) {
	fmt.Println(reverseBits(964176192))
	fmt.Println(reverseBits(43261596))

	fmt.Println(reverseBits(4294967293))
	fmt.Println(reverseBits(3221225471))
}

/*
【190. 颠倒二进制位】
颠倒给定的 32 位无符号整数的二进制位。

提示：
请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，
因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。

进阶:
如果多次调用这个函数，你将如何优化你的算法？

示例 1：
输入: 00000010100101000001111010011100
输出: 00111001011110000010100101000000
解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。

示例 2：
输入：11111111111111111111111111111101
输出：10111111111111111111111111111111
解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
     因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。

提示：
输入是一个长度为 32 的二进制字符串

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-bits
*/
func reverseBits(num uint32) uint32 {
	var r uint32
	for i := 0; i < 32; i++ {
		if (num>>i)&1 != 0 {
			r |= (1 << (31 - i))
		}
	}
	return r
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0074(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 1}}, 0))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

/*
【74. 搜索二维矩阵】
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。


示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10^4 <= matrix[i][j], target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return searchMatrix2(matrix, 0, len(matrix)*len(matrix[0]), target)
}

func searchMatrix2(matrix [][]int, from, to, target int) bool {
	if len(matrix) == 0 || from == to || from < 0 || to < 0 {
		return false
	}
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}
	var n = len(matrix[0])
	var index = from + (to-from)/2
	var row, col = index / n, index % n
	var f = matrix[row][col]
	switch {
	case f == target:
		return true
	case f < target:
		if index == from {
			return searchMatrix2(matrix, index+1, to, target)
		}
		return searchMatrix2(matrix, index, to, target)
	case f > target:
		if index == from {
			return searchMatrix2(matrix, from, index-1, target)
		}
		return searchMatrix2(matrix, from, index, target)
	}
	return false
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0090(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 3, 4, 5}))
	fmt.Println(subsetsWithDup([]int{9, 0, 3, 5, 7}))
	fmt.Println(subsetsWithDup([]int{1, 1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2}))
	fmt.Println(subsetsWithDup([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}

/*
【90. 子集 II】
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
*/
func subsetsWithDup(nums []int) [][]int {
	r := recursive0090(nums, len(nums))
	s := make([][]int, 0, len(r))
	for _, v := range r {
		var exist bool
		for j := 0; j < len(s); j++ {
			if arrayContains(v, s[j]) && arrayContains(s[j], v) {
				exist = true
				break
			}
		}
		if !exist {
			s = append(s, v)
		}
	}
	return s
}

// arrayContains 判断第一个数组包含第二个数组.
func arrayContains(a1, a2 []int) bool {
	if len(a1) < len(a2) {
		return false
	}
	for _, v2 := range a2 {
		var num int
		for _, v1 := range a1 {
			if v1 == v2 {
				num++
			}
		}
		if num == 0 {
			return false
		}
		for _, v1 := range a2 {
			if v1 == v2 {
				num--
			}
		}
		if num != 0 {
			return false
		}
	}
	return true
}

// recursive0090 从集合中取出元素个数在n以内的子集.
func recursive0090(nums []int, n int) [][]int {
	m := len(nums)
	if m == 0 || n == 0 {
		return [][]int{}
	}
	if n > m {
		n = m
	}
	switch n {
	case 0:
		return [][]int{}
	case 1:
		var r = make([][]int, m)
		for i, v := range nums {
			r[i] = []int{v}
		}
		return append(r, []int{})
	default:
		t := recursive0090(nums[1:], n-1)
		p := len(t)
		var r = make([][]int, 2*p)
		copy(r[:p], t)
		copy(r[p:2*p], t)
		for i := 0; i < p; i++ {
			r[i+p] = make([]int, len(r[i]))
			copy(r[i+p], r[i])
			r[i+p] = append(r[i+p], nums[0])
		}
		return r
	}
}
