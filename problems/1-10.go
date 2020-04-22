package problems

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//1. Two Sum
//最开始的想法是两个循环的暴力破解法，但是这样的时间复杂度是n²
//然后是利用map的两次循环，第一次循环把数组填入map，然后遍历数组，空间消耗会稍微大一点
//最后一种循环一次，如果map里面有组合值则返回，没有则放入map中
func TwoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for i, num := range nums {
		if value, ok := numMap[target-num]; ok {
			return []int{value, i}
		} else {
			numMap[num] = i
		}
	}
	return []int{}
}

//2. Add Two Numbers
//想转话为int的方法失败，遇到超长数组int64都无法解决
//准备直接用list直接相加进位来解决
//看了别人的代码 决定抛弃判断节点是首节点还是后续节点
//其实还可以把三个for循环合并
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode := new(ListNode)
	ptr := listNode
	var carryFlag int
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carryFlag
		carryFlag = val / 10
		val = val % 10

		node := &ListNode{
			Val: val,
		}
		ptr.Next = node
		ptr = node

		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + carryFlag
		carryFlag = val / 10
		val = val % 10

		node := &ListNode{
			Val: val,
		}
		ptr.Next = node
		ptr = node

		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + carryFlag
		carryFlag = val / 10
		val = val % 10

		node := &ListNode{
			Val: val,
		}
		ptr.Next = node
		ptr = node

		l2 = l2.Next
	}

	if carryFlag == 1 {
		node := &ListNode{
			Val: 1,
		}
		ptr.Next = node
		ptr = node
	}

	return listNode.Next
}

//3.Longest Substring Without Repeating Characters
//(1)用i和j两个指针在字符串上移动，保留最大长度变量和当前i到j的字符串，如果j的指针指向的字符在字符串内，判断字符是否最长字符，并移动i到j
//(2)上面的那个思路是有漏洞的，i应该移动到字符串内重复的字符后
//(3)判断字符的是否存在还能优化，不要去寻找，用数组来空间换时间
//func lengthOfLongestSubstring(s string) int {
//	var maxLength int
//	var i, j int
//	strLength := len(s)
//	length := 0
//	for i < strLength && j < strLength {
//		ok, index := isInSlice(s[i:j], s[j], i)
//		if ok {
//			i = index
//			length = j - index
//		} else {
//			j++
//			length++
//			if length > maxLength {
//				maxLength = length
//			}
//
//		}
//	}
//	return maxLength
//}
//
//func isInSlice(str string, element uint8, preIndex int) (bool, int) {
//	for i, char := range str {
//		if uint8(char) == element {
//			return true, i + preIndex + 1
//		}
//	}
//	return false, 0
//}
func lengthOfLongestSubstring(s string) int {
	character := [128]int{}
	maxLength, length := 0, 0
	start := 1
	for index, ch := range s {
		index = index + 1
		if character[ch] < start {
			length++
		} else if character[ch] >= start && character[ch] <= index {
			if length > maxLength {
				maxLength = length
			}
			start = character[ch] + 1
			length = index - start + 1
		}
		character[ch] = index
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}

//4.Median of Two Sorted Arrays
//时间复杂度要求为O(log(m+n)) 这点需要注意
//想法是把两个数组并到一起，然后利用快排的思想（偶数找到中间两位，基数找到中间位）这个的时间复杂度为0((n+m)log(n+m))
//审题不严，这里其实说了两个数组已经排序了
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	nums1 = append(nums1, nums2...)
//	length := len(nums1)
//	if length < 1 {
//		return 0
//	}
//
//	quickSort(&nums1, 0, length-1)
//	if length%2 == 0 {
//		return float64(nums1[length/2-1]+nums1[length/2]) / 2.0
//	} else {
//		return float64(nums1[length/2])
//	}
//}
//
//func quickSort(nums *[]int, start, end int) {
//	if start < end {
//		pos := position(nums, start, end)
//		quickSort(nums, start, pos-1)
//		quickSort(nums, pos+1, end)
//	}
//}

//func position(nums *[]int, start, end int) int {
//	target := (*nums)[end]
//	storeIndex := start
//	for i := start; i < end; i++ {
//		if (*nums)[i] < target {
//			(*nums)[i], (*nums)[storeIndex] = (*nums)[storeIndex], (*nums)[i]
//			storeIndex++
//		}
//	}
//	(*nums)[end], (*nums)[storeIndex] = (*nums)[storeIndex], (*nums)[end]
//	return storeIndex
//}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length == 0 {
		return 0.0
	}

	if length%2 == 1 {
		return findKth(nums1, nums2, length/2+1)
	} else {
		return (findKth(nums1, nums2, length/2) + findKth(nums1, nums2, length/2+1)) / 2.0
	}
}

func findKth(a []int, b []int, k int) float64 {
	if len(a) == 0 {
		return float64((b)[k-1])
	}
	if len(b) == 0 {
		return float64((a)[k-1])
	}
	if k == 1 {
		return math.Min(float64((a)[0]), float64((b)[0]))
	}
	var ak, bk int
	if k/2 > len(a) {
		ak = 1 << 31
	} else {
		ak = a[k/2-1]
	}

	if (k - k/2) > len(b) {
		bk = 1 << 31
	} else {
		bk = b[(k-k/2)-1]
	}

	if ak > bk {
		return findKth(a, b[(k-k/2):], k/2)
	} else if ak < bk {
		return findKth(a[k/2:], b, k-k/2)
	} else {
		return float64(ak)
	}
}

//5. Longest Palindromic Substring
//有种动态规划的方法，分奇偶，找到中间的数，然后向左右两边延伸，有个O（nk）的解法。（试了很久）
//func longestPalindrome(s string) string {
//	length := len(s)
//	if length == 0 {
//		return ""
//	}
//	if length == 1 {
//		return s
//	}
//
//	maxStr := string(s[0])
//	maxJ := 1
//	for i, _ := range s {
//		tempJ := 0
//		tempStr := ""
//		if i+1 < length && s[i] == s[i+1] {
//			for j := 1; j-1 <= i && i+j < length; j++ {
//				if s[i-j+1] == s[i+j] {
//					tempJ = 2 * j
//					tempStr = s[i-j+1 : i+j+1]
//				} else {
//					break
//				}
//			}
//		}
//
//		if tempJ > maxJ {
//			maxStr = tempStr
//			maxJ = tempJ
//		}
//
//		if i+1 < length && i-1 >= 0 && s[i-1] == s[i+1] {
//			for j := 1; j <= i && i+j < length; j++ {
//				if s[i-j] == s[i+j] {
//					tempJ = 2*j + 1
//					tempStr = s[i-j : i+j+1]
//				} else {
//					break
//				}
//			}
//		}
//
//		if tempJ > maxJ {
//			maxStr = tempStr
//			maxJ = tempJ
//		}
//	}
//	return maxStr
//}
//可以直接使用马拉车算法，时间复杂度是O（n）
func longestPalindrome(s string) string {
	t := "^#"
	for _, ch := range s {
		t = t + string(ch) + "#"
	}
	t += "$"
	length := len(t)
	id := 0
	mx := 0
	p := make([]int, length)
	max := 0
	for i, _ := range t {
		if i == 0 || i == length-1 {
			continue
		}
		if mx > i {
			j := 2*id - i
			if p[j] < mx-i {
				p[i] = p[j]
			} else {
				p[i] = mx - i
			}
		} else {
			p[i] = 1
		}
		for i-p[i] > 0 && i+p[i] < length {
			if t[i-p[i]] == t[i+p[i]] {
				p[i]++
			} else {
				break
			}
		}
		if p[i] > max {
			max = p[i]
			id = i
			mx = i + p[i]
		}
	}

	return s[(2*id-mx)/2 : mx/2-1]
}

//6. ZigZag Conversion
//1.第一个想法直接利用跳的形式获取字符，这样可能不行，因为Z字是不规则的
//2.z是以2numRows-2为一个循环的，一个一个的查找
//3.通过观察别人的答案，知道可以一横排一横排的找，不过字符串存在规律
//func convert(s string, numRows int) string {
//	ret := make([]byte, 0)
//	T := 2*numRows - 2
//	if T == 0 {
//		return s
//	}
//	for x := 1; x <= numRows; x++ {
//		for dT := 0; dT+x <= len(s); dT += T {
//
//			//fmt.Println("Append:", x,string(s[x - 1 + dT]))
//			ret = append(ret, s[x-1+dT])
//
//			if (dT+T-x+1 <= len(s)-1) && x != 1 && x != numRows {
//				//fmt.Println("AppendLast:", x, string(s[dT + T - x + 1]))
//				ret = append(ret, s[dT+T-x+1])
//			}
//		}
//	}
//	return string(ret)
//}
func Convert(s string, numRows int) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 || numRows == 1 {
		return s
	}
	strs := []string{}
	for i := 0; i < numRows; i++ {
		strs = append(strs, "")
	}
	returnStr := ""
	i := 0
	for ; len(s) > (i+1)*(2*numRows-2); i++ {
		Convert2n_2(&strs, s[i*(2*numRows-2):(i+1)*(2*numRows-2)], numRows)
	}
	Convert2n_2(&strs, s[i*(2*numRows-2):], numRows)

	for _, str := range strs {
		returnStr = returnStr + str
	}
	return returnStr
}

func Convert2n_2(strs *[]string, s string, numRows int) {
	for i := 0; i < len(s); i++ {
		if i < numRows {
			(*strs)[i] = (*strs)[i] + string(s[i])
		} else {
			(*strs)[2*numRows-2-i] = (*strs)[2*numRows-2-i] + string(s[i])
		}
	}
}

//7.Reverse Integer
//初看，感觉这道题很简单，但是题目有个i32需要注意
func reverse(x int) int {
	s := strconv.Itoa(x)
	t := ""
	tmp := ""
	for _, ch := range s {
		if ch == '-' {
			t = t + string(ch)
			continue
		}
		tmp = string(ch) + tmp
	}
	t = t + strings.TrimLeft(tmp, "0")

	//num, _ := strconv.ParseInt(t, 10, 32)
	num, _ := strconv.Atoi(t)
	if num > math.MaxInt32 {
		return 0
	}
	if num < math.MinInt32 {
		return 0
	}
	return num
}

//8.String to Integer(atoi)
//前面的空格要去掉（直接去掉前后的空白字符）,要去掉前后的加减号
//数字中额外的字符不需要处理
//如果第一个非空字符不是数字，或者字符串为空，不处理
//如果不能转换，返回0，如果能转换，返回最大值或者最小值
func myAtoi(str string) int {
	length := len(str)
	if length == 0 {
		return 0
	}
	var result int
	str1 := ""
	var isMinus bool
	str = strings.TrimSpace(str)

	for i, char := range str {
		if len(str1) == 0 {
			if (char == '-' || char == '+') && i+1 < length && isNum(int32(str[i+1])) {
				if char == '-' {
					isMinus = true
				}
			} else if (char == '-' || char == '+') && i+1 < length && !isNum(int32(str[i+1])) {
				return 0
			}
		}
		if isNum(char) {
			str1 = str1 + string(char)
		} else {
			if char != '+' && char != '-' {
				break
			}
		}
	}

	str1 = strings.TrimLeft(str1, "0")
	fmt.Println(str1)
	for _, ch := range str1 {
		result = result*10 + (int(ch) - 48)
		if isMinus && result >= math.MaxInt32+1 {
			return math.MinInt32
		} else if !isMinus && result >= math.MaxInt32 {
			return math.MaxInt32
		}

	}
	if isMinus {
		result = -1 * result
	}

	return result
}

func isNum(ch int32) bool {
	if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' ||
		ch == '8' || ch == '9' {
		return true
	}
	return false
}

//9.Palindrome Number
//如果是负数的话，就不是回文
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	y := 0
	for tmp > 0 {
		y = y*10 + tmp%10
		if y > math.MaxInt64 {
			return false
		}
		tmp /= 10
	}

	if x != y {
		return false
	}
	return true
}

//10.Regular Expression Matching
func isMatch(s string, p string) bool {
	if p == "" {
		return false
	}
	return true
}
