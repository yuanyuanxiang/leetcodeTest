package leetcodeTest

import (
	"fmt"
	"math"
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
	if len(arr) == 0 {
		return nil
	}
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

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0006(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("ABC", 2))
}

/*
【6.Z字形变换】
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N  0   4   8    4*i+0
A P L S I I G  1 3 5 7 9    4*i+1 4*i+3
Y   I   R      2   6   10   4*i+2
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);

示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N  0    6      12  6*i+0
A   L S  I G  1  5 7   11 13  6*i+1 6*i+5
Y A   H R     2 4  8 10   14  6*i+2 6*i+4
P     I       3    9      15  6*i+3

示例 3：
输入：s = "A", numRows = 1
输出："A"

提示：
1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
*/
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	n := len(s)
	m := (2*numRows - 2)
	buf := make([][]byte, numRows)
	buf[0] = make([]byte, 0, n/m+1)
	buf[numRows-1] = make([]byte, 0, n/m+1)
	for i := 1; i < numRows-1; i++ {
		buf[i] = make([]byte, 0, 2*n/m+1)
	}
	for i := 0; i < n; i++ {
		index := i % m
		if index >= numRows {
			index = m - index
		}
		buf[index] = append(buf[index], s[i])
	}
	var r = make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		r = append(r, buf[i]...)
	}

	return string(r)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0007(t *testing.T) {
	fmt.Println(reverse7(123))
	fmt.Println(reverse7(-123))
	fmt.Println(reverse7(120))
	fmt.Println(reverse7(0))
	fmt.Println(reverse7(math.MinInt32))
}

/*
【7.整数反转】
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0

提示：
-2^31 <= x <= 2^31 - 1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
*/
func reverse7(x int) int {
	var flag = x >= 0
	if !flag {
		x = -x
	}
	var arr = make([]int, 12)
	var num int
	for i, n := 0, x; n != 0; i++ {
		arr[i] = n % 10
		n /= 10
		num++
	}
	var k = 1
	var r = 0
	for i := 0; i < num; i++ {
		r += k * arr[num-1-i]
		k *= 10
	}
	if !flag {
		r = -r
		if r < math.MinInt32 {
			return 0
		}
		return r
	}
	if r > math.MaxInt32 {
		return 0
	}
	return r
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0008(t *testing.T) {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("+420"))
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("  4193 with words"))
	fmt.Println(myAtoi("0042"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi(".42"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("  0000000000012345678"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-9223372036854775808"))
	fmt.Println(myAtoi("18446744073709551617"))
	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	fmt.Println(myAtoi("-2147483648"))
	fmt.Println(myAtoi("+2147483648"))
	fmt.Println(myAtoi("-21474836480"))
	fmt.Println(myAtoi("+21474836470"))
}

/*
【8. 字符串转换整数 (atoi)】
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：
读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−2^31,  2^31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −2^31 ，
大于 2^31 − 1 的整数应该被固定为 2^31 − 1 。
返回整数作为最终结果。

注意：
本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
*/
func myAtoi(s string) int {
	n := len(s)
	var p int
	for ; p < n && s[p] == ' '; p++ {
	}
	if p == n {
		return 0
	}
	if s[p] != '+' && s[p] != '-' && !(s[p] >= '0' && s[p] <= '9') {
		return 0
	}
	var neg = s[p] == '-'
	if neg {
		p++
		if p == n || !(s[p] >= '0' && s[p] <= '9') {
			return 0
		}
	}
	if s[p] == '+' {
		p++
		if p == n || !(s[p] >= '0' && s[p] <= '9') {
			return 0
		}
	}
	for ; p < n && s[p] == '0'; p++ {
	}
	if p == n {
		return 0
	}
	var arr = make([]int, n-p)
	var i int
	for ; p < n && (s[p] >= '0' && s[p] <= '9'); i++ {
		arr[i] = int(s[p] - '0')
		p++
	}
	if i > 10 {
		if neg {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	var k = 1
	var r int
	for j := 0; j < i; j++ {
		r += k * arr[i-1-j]
		k *= 10
	}
	if neg {
		if -r < math.MinInt32 {
			return math.MinInt32
		}
		return -r
	}
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0009(t *testing.T) {
	fmt.Println("121:", isPalindrome(121))
	fmt.Println("-121:", isPalindrome(-121))
	fmt.Println("10:", isPalindrome(10))
	fmt.Println("-101:", isPalindrome(-101))
	fmt.Println("0:", isPalindrome(0))
}

/*
【9. 回文数】
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

示例 1：
输入：x = 121
输出：true

示例 2：
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3：
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。

示例 4：
输入：x = -101
输出：false

提示：
-2^31 <= x <= 2^31 - 1

进阶：你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var arr = make([]int, 12)
	i := 0
	for ; x != 0; i++ {
		arr[i] = x % 10
		x /= 10
	}
	for j := 0; j < i/2; j++ {
		if arr[j] != arr[i-1-j] {
			return false
		}
	}
	return true
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test0010(t *testing.T) {
	fmt.Println(isMatch("aaa", "ab*ac*a"))
	fmt.Println(isMatch("aaa", "a*a"))
	fmt.Println(isMatch("aaa", "a.a"))
	fmt.Println(isMatch("abcd", "d*"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("aa", ""))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("abcccc", ".*c"))
	fmt.Println(isMatch("abcccc", ".*cd"))
	fmt.Println(isMatch("abcccc", "a.*c"))
}

/*
【10. 正则表达式匹配】
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

示例 1：
输入：s = "aa" p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。

示例 2:
输入：s = "aa" p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3：
输入：s = "ab" p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4：
输入：s = "aab" p = "c*a*b"
输出：true
解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5：
输入：s = "mississippi" p = "mis*is*p*."
输出：false

提示：
0 <= s.length <= 20
0 <= p.length <= 30
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching

TODO 此题需要使用所谓“动态规划”的算法，本人尚不明白. 此题遗留.
*/
func isMatch(s string, p string) bool {
	return false
}
