package problems

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "测试1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "测试2",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 26,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "测试1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "测试2",
			args: args{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 8,
					},
				},
				l2: &ListNode{
					Val: 0,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
		{
			name: "测试3",
			args: args{
				l1: &ListNode{
					Val: 0,
				},
				l2: &ListNode{
					Val: 0,
				},
			},
			want: &ListNode{
				Val: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				s:       "",
				numRows: 3,
			},
			want: "",
		},
		{
			name: "测试2",
			args: args{
				s:       "P",
				numRows: 3,
			},
			want: "P",
		},
		{
			name: "测试3",
			args: args{
				s:       "PA",
				numRows: 4,
			},
			want: "PA",
		},
		{
			name: "测试4",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "测试5",
			args: args{
				s:       "AB",
				numRows: 1,
			},
			want: "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "测试2",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "测试3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "测试4",
			args: args{
				s: "c",
			},
			want: 1,
		},
		{
			name: "测试5",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_position(t *testing.T) {
//	type args struct {
//		nums  *[]int
//		start int
//		end   int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			name: "测试1",
//			args: args{
//				nums:  &[]int{1, 7, 6, 4},
//				start: 0,
//				end:   3,
//			},
//			want: 1,
//		},
//		{
//			name: "测试2",
//			args: args{
//				nums:  &[]int{1, 3, 2},
//				start: 0,
//				end:   2,
//			},
//			want: 1,
//		},
//		{
//			name: "测试3",
//			args: args{
//				nums:  &[]int{1, 1, 1, 1},
//				start: 1,
//				end:   3,
//			},
//			want: 1,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := position(tt.args.nums, tt.args.start, tt.args.end); got != tt.want {
//				t.Errorf("position() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_quickSort(t *testing.T) {
//	type args struct {
//		nums  *[]int
//		start int
//		end   int
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "测试1",
//			args: args{
//				nums:  &[]int{1, 1, 1, 1},
//				start: 0,
//				end:   3,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			quickSort(tt.args.nums, tt.args.start, tt.args.end)
//		})
//	}
//}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "测试1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "测试2",
			args: args{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			want: 2.5,
		},
		{
			name: "测试3",
			args: args{
				nums1: []int{1, 1},
				nums2: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "测试4",
			args: args{
				nums1: []int{1},
				nums2: []int{1},
			},
			want: 1.0,
		},
		{
			name: "测试5",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4, 5, 6},
			},
			want: 3.5,
		},
		{
			name: "测试6",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{1, 2},
			},
			want: 1.5,
		},
		{
			name: "测试7",
			args: args{
				nums1: []int{3, 4, 5, 6},
				nums2: []int{1, 2},
			},
			want: 3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "测试2",
			args: args{
				s: "babadada",
			},
			want: "adada",
		},
		{
			name: "测试3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "测试4",
			args: args{
				s: "tattarrattat",
			},
			want: "tattarrattat",
		},
		{
			name: "测试5",
			args: args{
				s: "ccc",
			},
			want: "ccc",
		},
		{
			name: "测试6",
			args: args{
				s: "cccc",
			},
			want: "cccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				x: 123,
			},
			want: 321,
		},
		{
			name: "2",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "3",
			args: args{
				x: 120,
			},
			want: 21,
		},
		{
			name: "4",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				str: "+1",
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				str: "",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				str: "+-2",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				str: "  -0012a42",
			},
			want: -12,
		},
		{
			name: "6",
			args: args{
				str: "9223372036854775809",
			},
			want: 2147483647,
		},
		{
			name: "7",
			args: args{
				str: "-2147483648",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				x: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
