package problems

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}},
		},
		{
			name: "2",
			args: args{
				candidates: []int{1},
				target:     2,
			},
			want: [][]int{},
		},
		{
			name: "2",
			args: args{
				candidates: []int{1},
				target:     1,
			},
			want: [][]int{{1}},
		},
		{
			name: "2",
			args: args{
				candidates: []int{1, 1},
				target:     2,
			},
			want: [][]int{{1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParentheses(t *testing.T) {
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
				s: "(()()",
			},
			want: 4,
		},
		{
			name: "测试2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "测试3",
			args: args{
				s: ")(",
			},
			want: 0,
		},
		{
			name: "测试4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "测试5",
			args: args{
				s: ")))))",
			},
			want: 0,
		},
		{
			name: "测试6",
			args: args{
				s: "()(())",
			},
			want: 6,
		},
		{
			name: "测试7",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
		{
			name: "测试8",
			args: args{
				s: ")()())()()(",
			},
			want: 4,
		},
		{
			name: "测试9",
			args: args{
				s: "(()))())(",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
