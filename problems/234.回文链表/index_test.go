package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *common.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{head: common.NewList([]interface{}{1, 2})},
			want: false,
		},
		{
			name: "test2",
			args: args{head: common.NewList([]interface{}{1, 2, 2, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
