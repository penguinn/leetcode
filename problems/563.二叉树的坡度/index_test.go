package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_findTilt(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2, 3})},
			want: 1,
		},
		{
			name: "test2",
			args: args{root: common.NewBinaryTree([]interface{}{1, 2})},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
