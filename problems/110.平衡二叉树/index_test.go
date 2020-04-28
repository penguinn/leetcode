package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{common.NewBinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7})},
			want: true,
		},
		{
			name: "test2",
			args: args{common.NewBinaryTree([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
