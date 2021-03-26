package leetcodeTest

import (
	"fmt"
	"strings"
	"testing"
)

func Test0001(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

/*
【1.两数之和】

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]


提示：
2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0002(t *testing.T) {
	fmt.Println(addTwoNumbers(NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})))
	fmt.Println(addTwoNumbers(NewListNode([]int{0}), NewListNode([]int{0})))
	fmt.Println(addTwoNumbers(NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), NewListNode([]int{9, 9, 9, 9})))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) *ListNode {
	var h = &ListNode{}
	var r = h
	for i, v := range arr {
		h.Val = v
		if i < len(arr)-1 {
			h.Next = &ListNode{}
			h = h.Next
		}
	}
	return r
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	return s
}

/*
【2.两数相加】
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：
每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p = &ListNode{}
	var head = p
	var prev = head
	for l1 != nil || l2 != nil {
		var a, b, val int
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		val = a + b
		p.Val += val
		if p.Val > 9 {
			p.Next = &ListNode{Val: 1}
			p.Val = p.Val - 10
		} else {
			p.Next = &ListNode{}
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		prev = p
		p = p.Next
	}
	if p.Val == 0 {
		prev.Next = nil
	}
	return head
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0002_0(t *testing.T) {
	fmt.Println(addTwoNumbers2([]int{1, 2, 3, 4}, []int{9, 2, 3, 6}))
	fmt.Println(addTwoNumbers2([]int{1, 2, 3, 4}, []int{9, 2, 3, 4}))
	fmt.Println(addTwoNumbers2([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}))
}

func inverse(arr []int) []int {
	n := len(arr)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[n-1-i] = arr[i]
	}
	return r
}

/*
【2.0两数相加】另外一个两数相加题目.
大数相加，给定两个大数，用数组表示，从高到低位，
依次存放在数组中，如 1234 表示为 [1,2,3,4]，对这个两个大数进行加和，返回结果也用数组表示
*/
func addTwoNumbers2(arr1, arr2 []int) []int {
	arr1 = inverse(arr1)
	arr2 = inverse(arr2)
	r := make([]int, 0)
	var flag int
	for i := 0; i < len(arr1) || i < len(arr2); i++ {
		var a, b int
		if i < len(arr1) {
			a = arr1[i]
		}
		if i < len(arr2) {
			b = arr2[i]
		}
		var sum = a + b + flag
		if sum > 9 {
			flag = 1
			r = append(r, sum-10)
		} else {
			flag = 0
			r = append(r, sum)
		}
	}
	if flag == 1 {
		r = append(r, 1)
	}
	return inverse(r)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0003(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}

/*
【3.无重复字符的最长子串】
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

示例 4:
输入: s = ""
输出: 0

提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	var max, m int
	for i, n := 0, len(s); i < n; i++ {
		m = n - i
		for j := i + 1; j < n; j++ {
			if strings.IndexByte(s[i:j], s[j]) != -1 {
				m = j - i
				break
			}
		}
		if m > max {
			max = m
		}
	}

	return max
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0004(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}

/*
【4.寻找两个正序数组的中位数】
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

示例 3：
输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000

示例 4：
输入：nums1 = [], nums2 = [1]
输出：1.00000

示例 5：
输入：nums1 = [2], nums2 = []
输出：2.00000

提示：
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m1 := len(nums1)
	m2 := len(nums2)
	var nums = make([]int, 0, m1+m2)
	for i, j := 0, 0; i < m1 || j < m2; {
		if i >= m1 {
			nums = append(nums, nums2[j])
			j++
			continue
		}
		if j >= m2 {
			nums = append(nums, nums1[i])
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
			continue
		}
		nums = append(nums, nums2[j])
		j++
	}
	if (m1+m2)%2 == 0 {
		return float64(nums[(m1+m2)/2-1]+nums[(m1+m2)/2]) / 2.0
	}
	return float64(nums[(m1+m2)/2])
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0005(t *testing.T) {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}

/*
【5.最长回文子串】
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

示例 3：
输入：s = "a"
输出："a"

示例 4：
输入：s = "ac"
输出："a"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
*/
func longestPalindrome(s string) string {
	var max int
	var r string
	for i, n := 0, len(s); i < n; i++ {
		for j := i + max; j < n; j++ {
			if t := s[i : j+1]; isHuiWen(t) {
				r = t
				max = j + 1 - i
			}
		}
	}
	return r
}

func TestReverse(t *testing.T) {
	fmt.Println(string(reverse([]byte(""))))
	fmt.Println(string(reverse([]byte("r"))))
	fmt.Println(string(reverse([]byte("hello"))))
	fmt.Println(string(reverse([]byte("test"))))
}

func reverse(s []byte) []byte {
	n := len(s)
	half := n / 2
	for i := 0; i < half; i++ {
		var t = s[i]
		s[i] = s[n-1-i]
		s[n-1-i] = t
	}
	return s
}

func TestHuiWen(t *testing.T) {
	fmt.Println(isHuiWen(""))
	fmt.Println(isHuiWen("ok"))
	fmt.Println(isHuiWen("oo"))
	fmt.Println(isHuiWen("oko"))
	fmt.Println(isHuiWen("okko"))
}

func isHuiWen(s string) bool {
	n := len(s)
	half := n / 2
	for i := 0; i < half; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
