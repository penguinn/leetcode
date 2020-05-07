package index

import (
	"testing"

	"github.com/penguinn/leetcode/common"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *common.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试1",
			args: args{root: common.NewBinaryTree([]interface{}{2, 1, 3})},
			want: true,
		},
		{
			name: "测试2",
			args: args{root: common.NewBinaryTree([]interface{}{10, 5, 15, nil, nil, 6, 20})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
